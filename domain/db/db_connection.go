package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DbConnect() *mongo.Client {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	uriEnv := os.Getenv("MONGO_PORT")
	uri := fmt.Sprintf("mongodb://localhost:%v", uriEnv)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_ = cancel

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	fmt.Println("Connect mongo")
	return client
}
