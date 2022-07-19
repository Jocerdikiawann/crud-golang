package userservices

import (
	usersdomain "belajar-golang-rest-api/models/domain/usersDomain"
	userrequests "belajar-golang-rest-api/models/requests/userRequests"
	userresponse "belajar-golang-rest-api/models/response/userResponse"
	userrepositories "belajar-golang-rest-api/repository/userRepositories"
	"belajar-golang-rest-api/utils"
	"context"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (c *UserServiceImpl) Create(ctx context.Context, req usersdomain.User) (usersdomain.User, error) {
	errvalidate := c.Validate.Struct(req)
	utils.IfErrorHandler(errvalidate)

	password := []byte(req.Password)
	hashedPassword, errPass := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	utils.IfErrorHandler(errPass)

	newdata := userrequests.UserRequest{
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Address:   req.Address,
	}

	result, err := c.UserRepo.Create(ctx, c.Db, newdata)
	return result, err
}

func (c *UserServiceImpl) GetUser(ctx context.Context, id string) (userresponse.UserResponse, error) {
	result, err := c.UserRepo.GetUser(ctx, c.Db, id)
	return result, err
}

func (c *UserServiceImpl) GetUsers(ctx context.Context) ([]userresponse.UserResponse, error) {
	result, err := c.UserRepo.GetUsers(ctx, c.Db)
	return result, err
}

func (c *UserServiceImpl) Update(ctx context.Context, request usersdomain.User) (userresponse.UserResponse, error) {

	password := []byte(request.Password)
	hashedPassword, errPass := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	utils.IfErrorHandler(errPass)

	objectId, _ := primitive.ObjectIDFromHex(request.Id)

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
	if result {
		getNewData, errNewData := c.UserRepo.GetUser(ctx, c.Db, request.Id)
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
