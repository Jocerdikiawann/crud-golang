package repository

import (
	"belajar-golang-rest-api/models/domain"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(ctx context.Context, db *mongo.Database, request domain.User) domain.User
	GetUser(ctx context.Context, db *mongo.Database, id string) domain.User
	GetUsers(ctx context.Context, db *mongo.Database) []domain.User
	Update(ctx context.Context, db *mongo.Database) domain.User
}
