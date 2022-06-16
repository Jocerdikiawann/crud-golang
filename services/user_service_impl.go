package services

import (
	"belajar-golang-rest-api/models/domain"
	"belajar-golang-rest-api/repository"
	"belajar-golang-rest-api/utils"
	"context"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	UserRepo repository.UserRepository
	Db       *mongo.Database
	Validate *validator.Validate
}

func NewUserService(repo repository.UserRepository, Db *mongo.Database, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepo: repo,
		Db:       Db,
		Validate: validate,
	}
}

func (c *UserServiceImpl) Create(ctx context.Context, request domain.User) domain.User {
	err := c.Validate.Struct(request)
	utils.IfErrorHandler(err)
	result := c.UserRepo.Create(ctx, c.Db, request)
	return result
}

func (c *UserServiceImpl) Getuser(ctx context.Context, id string) domain.User {
	result := c.UserRepo.GetUser(ctx, c.Db, id)
	return result
}
