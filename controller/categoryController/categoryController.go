package categorycontroller

import "github.com/gin-gonic/gin"

type CategoryController interface {
	Create(c *gin.Context)
	GetCategory(c *gin.Context)
	GetCategories(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
