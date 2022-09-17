package userrepositories

import (
	usersdomain "belajar-golang-rest-api/models/domain/usersDomain"
	userrequests "belajar-golang-rest-api/models/requests/userRequests"
	userresponse "belajar-golang-rest-api/models/response/userResponse"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	AuthSignIn(ctx context.Context, db *mongo.Database, request userrequests.AuthSignInRequest) (usersdomain.User, error)
	Create(ctx context.Context, db *mongo.Database, request userrequests.UserRequest) (userresponse.UserResponse, error)
	GetUser(ctx context.Context, db *mongo.Database, id string) (userresponse.UserResponse, error)
	GetUsers(ctx context.Context, db *mongo.Database) ([]userresponse.UserResponse, error)
	Update(ctx context.Context, db *mongo.Database, filter bson.M, request userrequests.UserRequest) (bool, error)
	Delete(ctx context.Context, db *mongo.Database, filter bson.M) (bool, error)
}
