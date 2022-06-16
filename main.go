// Package main provides ...
package main

import (
	"belajar-golang-rest-api/app"
	"belajar-golang-rest-api/controller"
	"belajar-golang-rest-api/repository"
	"belajar-golang-rest-api/services"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed load .env")
	}
}

func main() {
	dbUserName := os.Getenv("MONGO_USERNAME")
	dbPassword := os.Getenv("MONGO_PASSWORD")
	dbPort := os.Getenv("MONGO_PORT")
	dbName := os.Getenv("MONGO_DB_NAME")
	dbHost := os.Getenv("MONGO_HOST")

	Db := app.DbConnect(dbUserName, dbPassword, dbName, dbHost, dbPort)
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := services.NewUserService(userRepository, Db, validate)
	userController := controller.NewUserController(userService)

	router := app.NewRouter(userController)
	router.Run()
}
