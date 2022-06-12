// Package main provides ...
package main

import (
	"belajar-golang-rest-api/domain/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	db.DbConnect()
	c.JSON(http.StatusOK, "test aja")
}

func main() {
	router := gin.Default()
	router.GET("/", getUsers)
	router.Run(":8000")
}
