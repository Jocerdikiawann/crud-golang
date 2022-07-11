package app

import (
	"belajar-golang-rest-api/controller"
	"belajar-golang-rest-api/middlewares"
	"belajar-golang-rest-api/models/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func NewRouter(usercontroller controller.UserController) routes {
	r := routes{
		router: gin.New(),
	}

	r.router.Use(gin.Logger())
	r.router.Use(gin.Recovery())

	v1 := r.router.Group("/v1")
	{
		userRoutes := v1.Group("users")
		{
			userRoutes.POST("/", usercontroller.Create)
			userRoutes.GET("/:id", middlewares.MiddlewareAuth(), usercontroller.GetUser)
			userRoutes.GET("/", usercontroller.GetUsers)
			userRoutes.PUT("/:id", usercontroller.Update)
			userRoutes.DELETE("/:id", usercontroller.Delete)
		}
	}
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

func (r routes) Run(addr ...string) error {
	r.router.Use(gin.Recovery())
	return r.router.Run(addr...)
}
