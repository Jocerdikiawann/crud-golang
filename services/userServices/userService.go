package userservices

import (
	usersdomain "belajar-golang-rest-api/models/domain/usersDomain"
	userrequests "belajar-golang-rest-api/models/requests/userRequests"
	userresponse "belajar-golang-rest-api/models/response/userResponse"
	"context"
)

type UserService interface {
	AuthSignIn(ctx context.Context, req userrequests.AuthSignInRequest) (usersdomain.User, error)
	Create(ctx context.Context, req userrequests.UserRequest) (userresponse.UserResponse, error)
	GetUser(ctx context.Context, id string) (userresponse.UserResponse, error)
	GetUsers(ctx context.Context) ([]userresponse.UserResponse, error)
	Update(ctx context.Context, request userrequests.UserBody, id string) (userresponse.UserResponse, error)
	Delete(ctx context.Context, id string) ([]userresponse.UserResponse, error)
}
