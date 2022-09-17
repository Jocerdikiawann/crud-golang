package userrepositories

import (
	"belajar-golang-rest-api/middlewares"
	usersdomain "belajar-golang-rest-api/models/domain/usersDomain"
	userrequests "belajar-golang-rest-api/models/requests/userRequests"
	userresponse "belajar-golang-rest-api/models/response/userResponse"
	"belajar-golang-rest-api/utils"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct{}

// Init
func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) AuthSignIn(ctx context.Context, db *mongo.Database, request userrequests.AuthSignInRequest) (usersdomain.User, error) {
	var data usersdomain.User

	filter := bson.M{
		"email": request.Email,
	}

	result := db.Collection("user").FindOne(ctx, filter)

	errDecode := result.Decode(&data)
	utils.IfErrorHandler(errDecode)

	accessToken, errToken := middlewares.CreateToken(request.Email)
	utils.IfErrorHandler(errToken)

	data.AccessToken = accessToken
	return data, errDecode
}

func (repo *UserRepositoryImpl) Create(ctx context.Context, db *mongo.Database, request userrequests.UserRequest) (userresponse.UserResponse, error) {
	result, err := db.Collection("user").InsertOne(ctx, request)
	utils.IfErrorHandler(err)

	userdomain := userresponse.UserResponse{
		Id:        result.InsertedID.(primitive.ObjectID).Hex(),
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Address:   request.Address,
	}

	return userdomain, err
}

func (repo *UserRepositoryImpl) GetUser(ctx context.Context, db *mongo.Database, id string) (userresponse.UserResponse, error) {

	var data userrequests.UserRequest

	objId, err_ := primitive.ObjectIDFromHex(id)
	utils.IfErrorHandler(err_)

	filter := bson.M{
		"_id": objId,
	}

	result := db.Collection("user").FindOne(ctx, filter)

	err := result.Decode(&data)
	utils.IfErrorHandler(err)

	userData := userresponse.UserResponse{
		Id:          objId.Hex(),
		Email:       data.Email,
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Address:     data.Address,
		AccessToken: data.AccessToken,
	}

	return userData, err
}

func (repo *UserRepositoryImpl) GetUsers(ctx context.Context, db *mongo.Database) ([]userresponse.UserResponse, error) {

	var data []usersdomain.User
	var allData []userresponse.UserResponse

	filter := bson.M{}

	result, errDb := db.Collection("user").Find(ctx, filter)
	utils.IfErrorHandler(errDb)

	err := result.All(ctx, &data)
	utils.IfErrorHandler(err)

	for _, v := range data {
		tmp := userresponse.UserResponse{
			Id:          v.Id,
			Email:       v.Email,
			FirstName:   v.FirstName,
			LastName:    v.LastName,
			Address:     v.Address,
			AccessToken: v.AccessToken,
		}
		allData = append(allData, tmp)
	}

	return allData, errDb
}

func (repo *UserRepositoryImpl) Update(ctx context.Context, db *mongo.Database, filter bson.M, request userrequests.UserRequest) (bool, error) {

	_, err := db.Collection("user").UpdateOne(ctx, filter, bson.M{"$set": request})

	fmt.Println(err)
	if err != nil {
		return false, err
	}

	return true, err
}

func (repo *UserRepositoryImpl) Delete(ctx context.Context, db *mongo.Database, filter bson.M) (bool, error) {
	res, err := db.Collection("user").DeleteOne(ctx, filter)
	if res.DeletedCount == 0 {
		return false, err
	}
	return true, nil
}
