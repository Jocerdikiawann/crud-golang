// Package main provides ...
package main

import (
	"belajar-golang-rest-api/app"
	usercontroller "belajar-golang-rest-api/controller/userController"
	userrepositories "belajar-golang-rest-api/repository/userRepositories"
	userservices "belajar-golang-rest-api/services/userServices"
	"fmt"
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
	apiPort := fmt.Sprintf(":%v", os.Getenv("API_PORT"))

	Db := app.DbConnect(dbUserName, dbPassword, dbName, dbHost, dbPort)
	validate := validator.New()

	userRepository := userrepositories.NewUserRepository(validate)
	userService := userservices.NewUserService(userRepository, Db)
	userController := usercontroller.NewUserController(userService)

	// categoryRepository := categoryrepositories.NewCategoryRepository()
	// categoryService := categoryservices.NewCategoryService(categoryRepository, Db, validate)
	// categoryController := categorycontroller.NewCategoryController()

	router := app.NewRouter(userController)
	router.Run(apiPort)
}
