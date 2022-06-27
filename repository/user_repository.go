package repository

import (
	"belajar-golang-rest-api/models/requests"
	"belajar-golang-rest-api/models/response"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(ctx context.Context, db *mongo.Database, request requests.UserRequest) (response.UserResponse, error)
	GetUser(ctx context.Context, db *mongo.Database, id string) (response.UserResponse, error)
	GetUsers(ctx context.Context, db *mongo.Database) ([]response.UserResponse, error)
	Update(ctx context.Context, db *mongo.Database, filter bson.M, request requests.UserRequest) (bool, error)
	Delete(ctx context.Context, db *mongo.Database, filter bson.M) (bool, error)
}
