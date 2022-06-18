package repository

import (
	"belajar-golang-rest-api/models/domain"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(ctx context.Context, db *mongo.Database, request domain.User) (domain.User, error)
	GetUser(ctx context.Context, db *mongo.Database, id string) (domain.User, error)
	GetUsers(ctx context.Context, db *mongo.Database) ([]domain.User, error)
	Update(ctx context.Context, db *mongo.Database) (domain.User, error)
}
