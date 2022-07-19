package productsrepositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	Create(ctx context.Context, db *mongo.Database)
}
