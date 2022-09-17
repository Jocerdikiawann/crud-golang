package categoryservices

import (
	categorydomain "belajar-golang-rest-api/models/domain/categoryDomain"
	categoryrepositories "belajar-golang-rest-api/repository/categoryRepositories"
	"belajar-golang-rest-api/utils"
	"context"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryServiceImpl struct {
	CategoryRepo categoryrepositories.CategoryRepository
	Db           *mongo.Database
	Validate     *validator.Validate
}

func NewCategoryService(r categoryrepositories.CategoryRepository, db *mongo.Database, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepo: r,
		Db:           db,
		Validate:     validate,
	}
}

func (c *CategoryServiceImpl) Create(ctx context.Context, request categorydomain.CategoryRequest) (categorydomain.Category, error) {
	errValidate := c.Validate.Struct(request)
	utils.IfErrorHandler(errValidate)

	result, err := c.CategoryRepo.Create(ctx, c.Db, request)
	return result, err

}

func (c *CategoryServiceImpl) GetCategory(ctx context.Context, id string) (categorydomain.Category, error) {
	result, err := c.CategoryRepo.GetCategory(ctx, c.Db, id)
	return result, err
}

func (c *CategoryServiceImpl) GetCategories(ctx context.Context) ([]categorydomain.Category, error) {
	result, err := c.CategoryRepo.GetCategories(ctx, c.Db)
	return result, err
}

func (c *CategoryServiceImpl) Update(ctx context.Context, paramsId string, request categorydomain.CategoryRequest) (categorydomain.Category, error) {
	objId, _ := primitive.ObjectIDFromHex(paramsId)
	filter := bson.M{
		"_id": objId,
	}
	result, err := c.CategoryRepo.Update(ctx, c.Db, filter, request)
	if result {
		newData, errNewData := c.CategoryRepo.GetCategory(ctx, c.Db, paramsId)
		return newData, errNewData
	}
	return categorydomain.Category{}, err
}

func (c *CategoryServiceImpl) Delete(ctx context.Context, id string) ([]categorydomain.Category, error) {
	var newData []categorydomain.Category

	objId, errId := primitive.ObjectIDFromHex(id)
	utils.IfErrorHandler(errId)

	filter := bson.M{
		"_id": objId,
	}

	result, err := c.CategoryRepo.Delete(ctx, c.Db, filter)
	if result {
		newData, _ = c.CategoryRepo.GetCategories(ctx, c.Db)
		return newData, err
	}
	return newData, err
}
