package userrepositories

import (
	"belajar-golang-rest-api/models/user"
	"context"
)

type UserRepository interface {
	AuthSignIn(ctx context.Context, req user.AuthSignIn) (user.User, error)
	AuthSignUp(ctx context.Context, req user.AuthSignUp) (user.User, error)
	GetUser(ctx context.Context, id uint) (user.User, error)
	GetUsers(ctx context.Context) ([]user.User, error)
	Update(ctx context.Context, id uint, req user.UserUpdate) (user.User, error)
	Delete(ctx context.Context, id uint) error
}
