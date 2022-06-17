package services

import (
	"belajar-golang-rest-api/models/domain"
	"context"
)

type UserService interface {
	Create(ctx context.Context, request domain.User) domain.User
	GetUser(ctx context.Context, id string) domain.User
	GetUsers(ctx context.Context) []domain.User
	Update(ctx context.Context) domain.User
}
