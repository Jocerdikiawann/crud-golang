package services

import (
	"belajar-golang-rest-api/models/domain"
	"context"
)

type UserService interface {
	Create(ctx context.Context, request domain.User) (domain.User, error)
	GetUser(ctx context.Context, id string) (domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
	Update(ctx context.Context) (domain.User, error)
}
