package app

import (
	"belajar-golang-rest-api/controller"

	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func NewRouter(usercontroller controller.UserController) routes {
	r := routes{
		router: gin.Default(),
	}

	v1 := r.router.Group("/v1")
	r.addPing(usercontroller, v1)
	return r
}

func (r routes) addPing(usercontroller controller.UserController, rg *gin.RouterGroup) {
	ping := rg.Group("users")
	ping.POST("/", usercontroller.Create)
}

func (r routes) Run(addr ...string) error {
	return r.router.Run(addr...)
}
