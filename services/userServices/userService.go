package userservices

import (
	usersdomain "belajar-golang-rest-api/models/domain/usersDomain"
	userresponse "belajar-golang-rest-api/models/response/userResponse"
	"context"
)

type UserService interface {
	Create(ctx context.Context, req usersdomain.User) (usersdomain.User, error)
	GetUser(ctx context.Context, id string) (userresponse.UserResponse, error)
	GetUsers(ctx context.Context) ([]userresponse.UserResponse, error)
	Update(ctx context.Context, request usersdomain.User) (userresponse.UserResponse, error)
	Delete(ctx context.Context, id string) ([]userresponse.UserResponse, error)
}
