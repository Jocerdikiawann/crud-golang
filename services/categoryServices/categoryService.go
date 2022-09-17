package categoryservices

import (
	categorydomain "belajar-golang-rest-api/models/domain/categoryDomain"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, request categorydomain.CategoryRequest) (categorydomain.Category, error)
	GetCategory(ctx context.Context, id string) (categorydomain.Category, error)
	GetCategories(ctx context.Context) ([]categorydomain.Category, error)
	Update(ctx context.Context, paramsId string, request categorydomain.CategoryRequest) (categorydomain.Category, error)
	Delete(ctx context.Context, id string) ([]categorydomain.Category, error)
}
