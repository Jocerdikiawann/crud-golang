package categoryrepositories

import (
	categorydomain "belajar-golang-rest-api/models/domain/categoryDomain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository interface {
	Create(ctx context.Context, db *mongo.Database, request categorydomain.CategoryRequest) (categorydomain.Category, error)
	GetCategory(ctx context.Context, db *mongo.Database, id string) (categorydomain.Category, error)
	GetCategories(ctx context.Context, db *mongo.Database) ([]categorydomain.Category, error)
	Update(ctx context.Context, db *mongo.Database, filter bson.M, request categorydomain.CategoryRequest) (bool, error)
	Delete(ctx context.Context, db *mongo.Database, filter bson.M) (bool, error)
}
