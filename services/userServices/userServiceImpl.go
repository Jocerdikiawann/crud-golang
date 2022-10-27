package userservices

import (
	"belajar-golang-rest-api/models/response"
	"belajar-golang-rest-api/models/user"
	userrepositories "belajar-golang-rest-api/repository/userRepositories"
	"belajar-golang-rest-api/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepo userrepositories.UserRepository
	Db       *mongo.Database
}

func NewUserService(repo userrepositories.UserRepository, Db *mongo.Database) UserService {
	return &UserServiceImpl{
		UserRepo: repo,
		Db:       Db,
	}
}

func (c *UserServiceImpl) AuthSignIn(ctx *gin.Context) response.Response {

	var payload user.AuthSignIn

	err := ctx.BindJSON(&payload)

	userData, _ := c.UserRepo.AuthSignIn(ctx, c.Db, payload)

	errComparePassword := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(payload.Password))

	if errComparePassword != nil {
		return response.Response{
			StatusCode: http.StatusNotFound,
			Message:    "Wrong email or password",
			Data:       gin.H{},
		}
	}

	if err != nil {
		return response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       gin.H{},
		}
	}

	return response.Response{
		StatusCode: http.StatusOK,
		Message:    "ok",
		Data:       userData,
	}
}

func (c *UserServiceImpl) Create(ctx *gin.Context) response.Response {
	var req user.AuthSignUp

	errJson := ctx.BindJSON(&req)

	password := []byte(req.Password)
	hashedPassword, errPass := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	newdata := user.AuthSignUp{
		Email:     req.Email,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Address:   req.Address,
	}

	result, errData := c.UserRepo.Create(ctx, c.Db, newdata)

	if errJson != nil {
		return response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    errJson.Error(),
			Data:       gin.H{},
		}
	}

	if errPass != nil {
		return response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    errPass.Error(),
			Data:       gin.H{},
		}
	}

	if errData != nil {
		return response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    errData.Error(),
			Data:       gin.H{},
		}
	}

	return response.Response{
		StatusCode: http.StatusCreated,
		Message:    "ok",
		Data:       result,
	}
}

func (c *UserServiceImpl) GetUser(ctx *gin.Context) response.Response {

	id := ctx.Param("id")

	result, err := c.UserRepo.GetUser(ctx, c.Db, id)

	if err != nil {
		return response.Response{
			StatusCode: http.StatusNotFound,
			Message:    "Account not found",
			Data:       gin.H{},
		}
	}

	return response.Response{
		StatusCode: http.StatusOK,
		Message:    "ok",
		Data:       result,
	}
}

func (c *UserServiceImpl) GetUsers(ctx *gin.Context) response.Response {
	result, err := c.UserRepo.GetUsers(ctx, c.Db)

	if err != nil {
		return response.Response{
			StatusCode: http.StatusNotFound,
			Message:    "Account not found",
			Data:       gin.H{},
		}
	}

	return response.Response{
		StatusCode: http.StatusOK,
		Message:    "ok",
		Data:       result,
	}
}

func (c *UserServiceImpl) Update(ctx *gin.Context) response.Response {

	id := ctx.Param("id")

	var req user.AuthSignUp

	errJson := ctx.BindJSON(&req)

	password := []byte(req.Password)
	hashedPassword, errPass := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	utils.IfErrorHandler(errPass)

	objectId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{
		"_id": objectId,
	}
	updatedData := userrequests.UserRequest{
		Email:     request.Email,
		Password:  hashedPassword,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Address:   request.Address,
	}
	result, err := c.UserRepo.Update(ctx, c.Db, filter, updatedData)

	if errJson != nil {
		return response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    errJson.Error(),
			Data:       gin.H{},
		}
	}
	
	if result {
		getNewData, errNewData := c.UserRepo.GetUser(ctx, c.Db, id)
		return getNewData, errNewData
	}

	return userresponse.UserResponse{}, err
}

func (c *UserServiceImpl) Delete(ctx context.Context, id string) ([]userresponse.UserResponse, error) {

	var newDataUser []userresponse.UserResponse

	objId, errId := primitive.ObjectIDFromHex(id)
	utils.IfErrorHandler(errId)

	filter := bson.M{
		"_id": objId,
	}

	result, err := c.UserRepo.Delete(ctx, c.Db, filter)
	if result {
		newDataUser, _ = c.UserRepo.GetUsers(ctx, c.Db)
		return newDataUser, err
	}
	return newDataUser, err
}
