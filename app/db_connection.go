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
	uri := fmt.Sprintf("mongodb://localhost:%v", portDb)

	fmt.Println("mongo url: ", uri)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	_ = cancel

	//credential := options.Credential{
	//	AuthSource: "admin",
	//	Username:   usernameDb,
	//	Password:   passwordDb,
	//}

	clientOptions := options.Client().ApplyURI(uri)
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
