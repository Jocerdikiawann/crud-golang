package usercontroller

import "github.com/gin-gonic/gin"

type UserController interface {
	AuthSignIn(c *gin.Context)
	Create(c *gin.Context)
	GetUser(c *gin.Context)
	GetUsers(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
