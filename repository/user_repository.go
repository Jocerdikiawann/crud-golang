package repository

import (
	"belajar-golang-rest-api/models/domain"
	"belajar-golang-rest-api/models/response"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(ctx context.Context, db *mongo.Database, request domain.User) (response.UserResponse, []error)
	GetUser(ctx context.Context, db *mongo.Database, id string) (response.UserResponse, []error)
	GetUsers(ctx context.Context, db *mongo.Database) ([]response.UserResponse, []error)
	Update(ctx context.Context, db *mongo.Database) (response.UserResponse, []error)
}
