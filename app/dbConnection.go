package app

import (
	"belajar-golang-rest-api/utils"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DbConnect(usernameDb, passwordDb, nameDb, hostDb, portDb string) *mongo.Database {
	uri := fmt.Sprintf("mongodb://%s:%v", hostDb, portDb)

	credential := options.Credential{
		Username: usernameDb,
		Password: passwordDb,
	}

	fmt.Println("mongo url: ", uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri).SetAuth(credential).SetMaxPoolSize(100).SetMinPoolSize(20)
	client, err := mongo.NewClient(clientOptions)
	utils.IfErrorHandler(err)

	err = client.Connect(ctx)
	utils.IfErrorHandler(err)

	err = client.Ping(ctx, readpref.Primary())
	utils.IfErrorHandler(err)

	db := client.Database(nameDb)

	fmt.Println("Connect mongo")
	return db
}
