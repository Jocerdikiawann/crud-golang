package categoryrepositories

import (
	categorydomain "belajar-golang-rest-api/models/domain/categoryDomain"
	"belajar-golang-rest-api/utils"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (r *CategoryRepositoryImpl) Create(ctx context.Context, db *mongo.Database, request categorydomain.CategoryRequest) (categorydomain.Category, error) {
	result, err := db.Collection("category").InsertOne(ctx, request)
	utils.IfErrorHandler(err)

	categoryDomain := categorydomain.Category{
		Id:             result.InsertedID.(primitive.ObjectID).Hex(),
		NameOfCategory: request.NameOfCategory,
	}

	return categoryDomain, err
}

func (r *CategoryRepositoryImpl) GetCategory(ctx context.Context, db *mongo.Database, id string) (categorydomain.Category, error) {
	var data categorydomain.Category

	objId, errObjId := primitive.ObjectIDFromHex(id)
	utils.IfErrorHandler(errObjId)

	filter := bson.M{
		"_id": objId,
	}

	err := db.Collection("category").FindOne(ctx, filter).Decode(&data)
	utils.IfErrorHandler(err)

	catagory := categorydomain.Category{
		Id:             objId.Hex(),
		NameOfCategory: data.NameOfCategory,
	}

	return catagory, err
}

func (r *CategoryRepositoryImpl) GetCategories(ctx context.Context, db *mongo.Database) ([]categorydomain.Category, error) {
	var data []categorydomain.Category

	result, errFilter := db.Collection("category").Find(ctx, bson.M{})
	utils.IfErrorHandler(errFilter)
	err := result.All(ctx, &data)
	utils.IfErrorHandler(err)

	return data, err
}

func (r *CategoryRepositoryImpl) Update(ctx context.Context, db *mongo.Database, filter bson.M, request categorydomain.CategoryRequest) (bool, error) {
	_, err := db.Collection("category").UpdateOne(ctx, filter, bson.M{"$set": request})
	utils.IfErrorHandler(err)

	if err != nil {
		return false, err
	}
	return true, err
}

func (r *CategoryRepositoryImpl) Delete(ctx context.Context, db *mongo.Database, filter bson.M) (bool, error) {
	result, err := db.Collection("category").DeleteOne(ctx, filter)
	if result.DeletedCount == 0 {
		return false, err
	}
	return true, err
}
