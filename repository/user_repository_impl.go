package repository

import (
	"belajar-golang-rest-api/models/domain"
	"belajar-golang-rest-api/utils"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct{}

//Init
func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) Create(ctx context.Context, db *mongo.Database, request domain.User) domain.User {
	result, err := db.Collection("user").InsertOne(ctx, request)
	utils.IfErrorHandler(err)

	return domain.User{
		Id:        result.InsertedID.(primitive.ObjectID).Hex(),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Address:   request.Address,
	}
}

func (repo *UserRepositoryImpl) GetUser(ctx context.Context, db *mongo.Database, id string) domain.User {
	data := domain.User{}
	objId, err := primitive.ObjectIDFromHex(id)
	utils.IfErrorHandler(err)
	filter := bson.M{
		"_id": objId,
	}
	result := db.Collection("user").FindOne(ctx, filter)
	result.Decode(&data)
	return data
}

func (repo *UserRepositoryImpl) GetUsers(ctx context.Context, db *mongo.Database) []domain.User {
	var data []domain.User
	filter := bson.M{}
	result, err := db.Collection("user").Find(ctx, filter)
	utils.IfErrorHandler(err)
	error := result.All(ctx, &data)
	utils.IfErrorHandler(error)
	return data
}

func (repo *UserRepositoryImpl) Update(ctx context.Context, db *mongo.Database) domain.User {
	return domain.User{}
}
