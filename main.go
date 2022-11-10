// Package main provides ...
package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed load .env")
	}
}

func main() {
	// dbUserName := os.Getenv("MONGO_USERNAME")
	// dbPassword := os.Getenv("MONGO_PASSWORD")
	// dbPort := os.Getenv("MONGO_PORT")
	// dbName := os.Getenv("MONGO_DB_NAME")
	// dbHost := os.Getenv("MONGO_HOST")
	// apiPort := fmt.Sprintf(":%v", os.Getenv("API_PORT"))

	// Db := app.DbConnect(dbUserName, dbPassword, dbName, dbHost, dbPort)
	// validate := validator.New()

	// userRepository := userrepositories.NewUserRepository()
	// userService := userservices.NewUserService(userRepository, Db, validate)
	// userController := usercontroller.NewUserController(userService)

	// roleRepository := rolerepositories.NewRoleRepository()
	// roleServices := roleservices.NewRoleService(roleRepository, Db, validate)
	// roleController := rolecontroller.NewRoleController(roleServices)
	// categoryRepository := categoryrepositories.NewCategoryRepository()
	// categoryService := categoryservices.NewCategoryService(categoryRepository, Db, validate)
	// categoryController := categorycontroller.NewCategoryController()

	// router := app.NewRouter(userController, roleController)
	// router.Run(apiPort)
}
