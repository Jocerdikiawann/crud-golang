// Package main provides ...
package main

import (
	"belajar-golang-rest-api/app"
	categorycontroller "belajar-golang-rest-api/controller/categoryController"
	usercontroller "belajar-golang-rest-api/controller/userController"
	categoryrepositories "belajar-golang-rest-api/repository/categoryRepositories"
	userrepositories "belajar-golang-rest-api/repository/userRepositories"
	categoryservices "belajar-golang-rest-api/services/categoryServices"
	userservices "belajar-golang-rest-api/services/userServices"
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

	userRepository := userrepositories.NewUserRepository()
	userService := userservices.NewUserService(userRepository, Db, validate)
	userController := usercontroller.NewUserController(userService)

	categoryRepository := categoryrepositories.NewCategoryRepository()
	categoryService := categoryservices.NewCategoryService(categoryRepository, Db, validate)
	categoryController := categorycontroller.NewCategoryController(categoryService)

	router := app.NewRouter(userController, categoryController)
	router.Run(":8000")
}
