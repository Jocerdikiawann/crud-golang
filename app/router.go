package app

import (
	categorycontroller "belajar-golang-rest-api/controller/categoryController"
	usercontroller "belajar-golang-rest-api/controller/userController"
	"belajar-golang-rest-api/middlewares"
	"belajar-golang-rest-api/models/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func NewRouter(user usercontroller.UserController, category categorycontroller.CategoryController) routes {
	r := routes{
		router: gin.New(),
	}
	r.router.RedirectTrailingSlash = false

	r.router.Use(gin.Logger())
	r.router.Use(gin.Recovery())

	v1 := r.router.Group("/v1")
	{
		userRoutes := v1.Group("users")
		{
			userRoutes.POST("/auth", user.AuthSignIn)
			userRoutes.POST("/", user.Create)
			userRoutes.GET("/:id", middlewares.MiddlewareAuth(), user.GetUser)
			userRoutes.GET("/", middlewares.MiddlewareAuth(), user.GetUsers)
			userRoutes.PUT("/:id", middlewares.MiddlewareAuth(), user.Update)
			userRoutes.DELETE("/:id", middlewares.MiddlewareAuth(), user.Delete)
		}
		categoryRoutes := v1.Group("category").Use(middlewares.MiddlewareAuth())
		{
			categoryRoutes.POST("/", category.Create)
			categoryRoutes.GET("/:id", category.GetCategory)
			categoryRoutes.GET("/", category.GetCategories)
			categoryRoutes.PUT("/:id", category.Update)
			categoryRoutes.DELETE("/:id", category.Delete)
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
