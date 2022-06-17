package app

import (
	"belajar-golang-rest-api/controller"
	"belajar-golang-rest-api/models/response"
	"net/http"

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

	r.routeNotFound()

	return r
}

func (r routes) routeNotFound() {
	r.router.NoRoute(func(c *gin.Context) {
		res := response.WebResponse{
			StatusCode: http.StatusNotFound,
			Message:    "route not found",
			Data:       gin.H{},
		}
		c.IndentedJSON(res.StatusCode, res)
	})
}

func (r routes) addPing(usercontroller controller.UserController, rg *gin.RouterGroup) {
	ping := rg.Group("users")
	ping.POST("/", usercontroller.Create)
	ping.GET("/:id", usercontroller.GetUser)
	ping.GET("/", usercontroller.GetUsers)
}

func (r routes) Run(addr ...string) error {
	r.router.Use(gin.Recovery())
	return r.router.Run(addr...)
}
