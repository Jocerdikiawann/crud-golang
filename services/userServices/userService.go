package userservices

import (
	"belajar-golang-rest-api/models/response"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	AuthSignIn(ctx *gin.Context) response.Response
	Create(ctx *gin.Context) response.Response
	GetUser(ctx *gin.Context) response.Response
	GetUsers(ctx *gin.Context) response.Response
	Update(ctx *gin.Context) response.Response
	Delete(ctx *gin.Context) response.Response
}
