package services

import (
	usersdomain "belajar-golang-rest-api/models/domain/users_domain"
	"belajar-golang-rest-api/models/response"
	"context"
)

type UserService interface {
	Create(ctx context.Context, req usersdomain.User) (usersdomain.User, error)
	GetUser(ctx context.Context, id string) (response.UserResponse, error)
	GetUsers(ctx context.Context) ([]response.UserResponse, error)
	Update(ctx context.Context, request usersdomain.User) (response.UserResponse, error)
	Delete(ctx context.Context, id string) ([]response.UserResponse, error)
}
