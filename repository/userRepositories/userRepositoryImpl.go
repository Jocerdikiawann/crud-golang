package userrepositories

import (
	"belajar-golang-rest-api/middlewares"
	"belajar-golang-rest-api/models/user"
	"belajar-golang-rest-api/utils"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
}

// Init
func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) AuthSignIn(ctx context.Context, db *mongo.Database, req user.AuthSignIn) (user.User, error) {
	var data user.User

	filter := bson.M{
		"email": req.Email,
	}

	result := db.Collection("user").FindOne(ctx, filter)

	errDecode := result.Decode(&data)
	utils.IfErrorHandler(errDecode)

	accessToken, errToken := middlewares.CreateToken(data.Id)
	utils.IfErrorHandler(errToken)

	data.AccessToken = accessToken
	return data, errDecode
}

func (repo *UserRepositoryImpl) Create(ctx context.Context, db *mongo.Database, req user.AuthSignUp) (user.User, error) {

	coll := db.Collection("user")

	result, err := coll.InsertOne(ctx, req)
	utils.Error.Println(err)

	if err != nil {
		return user.User{}, err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	accessToken, errToken := middlewares.CreateToken(id)
	utils.IfErrorHandler(errToken)

	createdUser := user.User{
		Id:          id,
		Email:       req.Email,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Address:     req.Address,
		AccessToken: accessToken,
	}

	return createdUser, err
}

func (repo *UserRepositoryImpl) GetUser(ctx context.Context, db *mongo.Database, id string) (user.User, error) {

	var data user.User

	objId, errId := primitive.ObjectIDFromHex(id)
	utils.IfErrorHandler(errId)

	filter := bson.M{
		"_id": objId,
	}

	result := db.Collection("user").FindOne(ctx, filter)

	err := result.Decode(&data)
	utils.IfErrorHandler(err)

	userData := user.User{
		Id:        objId.Hex(),
		Email:     data.Email,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Address:   data.Address,
	}

	return userData, err
}

func (repo *UserRepositoryImpl) GetUsers(ctx context.Context, db *mongo.Database) ([]user.User, error) {

	var data []user.User
	var newData []user.User

	filter := bson.M{}

	result, errDb := db.Collection("user").Find(ctx, filter)
	utils.IfErrorHandler(errDb)

	err := result.All(ctx, &data)
	utils.IfErrorHandler(err)

	for _, v := range data {
		tmp := user.User{
			Id:        v.Id,
			Email:     v.Email,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Address:   v.Address,
		}
		newData = append(newData, tmp)
	}

	return newData, errDb
}

func (repo *UserRepositoryImpl) Update(ctx context.Context, db *mongo.Database, id string, req user.AuthSignUp) (bool, error) {

	objId, errId := primitive.ObjectIDFromHex(id)
	utils.IfErrorHandler(errId)

	_, err := db.Collection("user").UpdateOne(
		ctx,
		bson.M{"_id": objId},
		bson.M{"$set": req},
	)

	if err != nil {
		return false, err
	}

	return true, err
}

func (repo *UserRepositoryImpl) Delete(ctx context.Context, db *mongo.Database, id string) (bool, error) {
	objId, errId := primitive.ObjectIDFromHex(id)
	utils.IfErrorHandler(errId)

	res, err := db.Collection("user").DeleteOne(ctx, bson.M{"_id": objId})
	if res.DeletedCount == 0 {
		return false, err
	}
	return true, nil
}
