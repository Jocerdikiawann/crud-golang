package services

import (
	"belajar-golang-rest-api/models/domain"
	"context"
)

type UserService interface {
	Create(ctx context.Context, request domain.User) domain.User
	Getuser(ctx context.Context, id string) domain.User
}
