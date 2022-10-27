package userrepositories

import (
	"belajar-golang-rest-api/models/user"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	AuthSignIn(ctx context.Context, db *mongo.Database, req user.AuthSignIn) (user.User, error)
	Create(ctx context.Context, db *mongo.Database, req user.AuthSignUp) (user.User, error)
	GetUser(ctx context.Context, db *mongo.Database, id string) (user.User, error)
	GetUsers(ctx context.Context, db *mongo.Database) ([]user.User, error)
	Update(ctx context.Context, db *mongo.Database, id string, req user.AuthSignUp) (bool, error)
	Delete(ctx context.Context, db *mongo.Database, id string) (bool, error)
}
