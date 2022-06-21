package services

import (
	"belajar-golang-rest-api/models/domain"
	"belajar-golang-rest-api/models/response"
	"context"
)

type UserService interface {
	Create(ctx context.Context, request domain.User) (response.UserResponse, []error)
	GetUser(ctx context.Context, id string) (response.UserResponse, []error)
	GetUsers(ctx context.Context) ([]response.UserResponse, []error)
	Update(ctx context.Context) (response.UserResponse, []error)
}
