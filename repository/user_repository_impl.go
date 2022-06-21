package repository

import (
	"belajar-golang-rest-api/models/domain"
	"belajar-golang-rest-api/models/requests"
	"belajar-golang-rest-api/models/response"
	"belajar-golang-rest-api/utils"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImpl struct{}

//Init
func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) Create(ctx context.Context, db *mongo.Database, request domain.User) (response.UserResponse, []error) {

	var errs []error

	password := []byte(request.Password)
	hashedPassword, errPass := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	fmt.Println("hash", hashedPassword)
	utils.IfErrorHandler(errPass)

	requestData := requests.UserRequest{
		Email:     request.Email,
		Password:  hashedPassword,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Address:   request.LastName,
	}

	result, err := db.Collection("user").InsertOne(ctx, requestData)
	utils.IfErrorHandler(err)

	if errPass != nil || err != nil {
		errs = append(errs, errPass, err)
	}

	return response.UserResponse{
		Id:        result.InsertedID.(primitive.ObjectID).Hex(),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Address:   request.Address,
	}, errs
}

func (repo *UserRepositoryImpl) GetUser(ctx context.Context, db *mongo.Database, id string) (response.UserResponse, []error) {

	var errs []error
	data := requests.UserRequest{}

	objId, err := primitive.ObjectIDFromHex(id)
	utils.IfErrorHandler(err)

	filter := bson.M{
		"_id": objId,
	}

	result := db.Collection("user").FindOne(ctx, filter)

	errD := result.Decode(&data)
	utils.IfErrorHandler(errD)
	if err != nil || errD != nil {
		errs = append(errs, err, errD)
	}

	userData := response.UserResponse{
		Id:        objId.Hex(),
		Email:     data.Email,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Address:   data.Address,
	}

	return userData, errs
}

func (repo *UserRepositoryImpl) GetUsers(ctx context.Context, db *mongo.Database) ([]response.UserResponse, []error) {

	var errs []error
	var data []domain.User
	var allData []response.UserResponse

	filter := bson.M{}

	result, err := db.Collection("user").Find(ctx, filter)
	utils.IfErrorHandler(err)

	error_ := result.All(ctx, &data)
	utils.IfErrorHandler(error_)

	for _, v := range data {
		tmp := response.UserResponse{
			Id:        v.Id,
			Email:     v.Email,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Address:   v.Address,
		}
		allData = append(allData, tmp)
	}

	if err != nil || error_ != nil {
		errs = append(errs, err, error_)
	}

	return allData, errs
}

func (repo *UserRepositoryImpl) Update(ctx context.Context, db *mongo.Database) (response.UserResponse, []error) {
	return response.UserResponse{}, nil
}
