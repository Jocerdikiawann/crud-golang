package roleservices

import (
	"belajar-golang-rest-api/models/response"

	"github.com/gin-gonic/gin"
)

type RoleService interface {
	Create(ctx *gin.Context) response.Response
}
