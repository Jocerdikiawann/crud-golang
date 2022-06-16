package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Create(c *gin.Context)
	Getuser(c *gin.Context)
}
