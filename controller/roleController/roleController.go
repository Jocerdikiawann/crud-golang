package rolecontroller

import "github.com/gin-gonic/gin"

type RoleController interface {
	Create(c *gin.Context)
}