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

func (c *UserServiceImpl) Create(ctx context.Context, request domain.User) (domain.User, error) {
	e := c.Validate.Struct(request)
	utils.IfErrorHandler(e)
	result, err := c.UserRepo.Create(ctx, c.Db, request)
	return result, err
}

func (c *UserServiceImpl) GetUser(ctx context.Context, id string) (domain.User, error) {
	result, err := c.UserRepo.GetUser(ctx, c.Db, id)
	return result, err
}

func (c *UserServiceImpl) GetUsers(ctx context.Context) ([]domain.User, error) {
	result, err := c.UserRepo.GetUsers(ctx, c.Db)
	return result, err
}

func (c *UserServiceImpl) Update(ctx context.Context) (domain.User, error) {
	result, err := c.UserRepo.Update(ctx, c.Db)
	return result, err
}
