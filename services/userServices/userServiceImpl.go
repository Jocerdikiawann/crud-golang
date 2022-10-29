package userservices

import (
	"belajar-golang-rest-api/models/response"
	"belajar-golang-rest-api/models/user"
	userrepositories "belajar-golang-rest-api/repository/userRepositories"
	"belajar-golang-rest-api/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepo userrepositories.UserRepository
	Db       *mongo.Database
	Validate *validator.Validate
}

func NewUserService(repo userrepositories.UserRepository, Db *mongo.Database, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepo: repo,
		Db:       Db,
		Validate: validate,
	}
}

func (c *UserServiceImpl) AuthSignIn(ctx *gin.Context) response.Response {

	var payload user.AuthSignIn

	err := ctx.BindJSON(&payload)

	errValidate := c.Validate.Struct(payload)
	utils.IfErrorHandler(errValidate)

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

	errValidate := c.Validate.Struct(req)
	utils.IfErrorHandler(errValidate)

	hashedPassword, errPass := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	req.Password = string(hashedPassword)

	result, errData := c.UserRepo.Create(ctx, c.Db, req)

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

	id := string(ctx.MustGet("_id").(string))

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

	id := string(ctx.MustGet("_id").(string))

	var req user.AuthSignUp

	errJson := ctx.BindJSON(&req)

	password := []byte(req.Password)
	hashedPassword, errPass := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	req.Password = string(hashedPassword)

	fmt.Printf("Password %v", req.Password)

	result, err := c.UserRepo.Update(ctx, c.Db, id, req)

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

	if err != nil {
		return response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       gin.H{},
		}
	}

	if result {
		getNewData, errNewData := c.UserRepo.GetUser(ctx, c.Db, id)
		if errNewData != nil {
			return response.Response{
				StatusCode: http.StatusBadRequest,
				Message:    errNewData.Error(),
				Data:       gin.H{},
			}
		}
		return response.Response{
			StatusCode: http.StatusOK,
			Message:    "ok",
			Data:       getNewData,
		}
	}
	return response.Response{
		StatusCode: http.StatusOK,
		Message:    "ok",
		Data:       []string{},
	}
}

func (c *UserServiceImpl) Delete(ctx *gin.Context) response.Response {

	id := string(ctx.MustGet("_id").(string))

	result, err := c.UserRepo.Delete(ctx, c.Db, id)

	if err != nil || !result {
		return response.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       gin.H{},
		}
	}

	res, _ := c.UserRepo.GetUsers(ctx, c.Db)

	return response.Response{
		StatusCode: http.StatusOK,
		Message:    "ok",
		Data:       res,
	}
}
