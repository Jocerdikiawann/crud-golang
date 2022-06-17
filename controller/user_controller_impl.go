package controller

import (
	"belajar-golang-rest-api/models/domain"
	"belajar-golang-rest-api/models/response"
	"belajar-golang-rest-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	service services.UserService
}

func NewUserController(services services.UserService) UserController {
	return &UserControllerImpl{
		services,
	}
}

func (controller *UserControllerImpl) Create(c *gin.Context) {
	var body domain.User
	if err := c.BindJSON(&body); err != nil {
		return
	}

	data := controller.service.Create(c.Request.Context(), body)

	responseJson := response.WebResponse{
		StatusCode: http.StatusCreated,
		Message:    "ok",
		Data:       data,
	}

	c.IndentedJSON(http.StatusCreated, responseJson)
}

func (controller *UserControllerImpl) GetUser(c *gin.Context) {
	id := c.Param("id")
	data := controller.service.GetUser(c.Request.Context(), id)

	responseJson := response.WebResponse{
		StatusCode: http.StatusOK,
		Message:    "ok",
		Data:       data,
	}
	c.IndentedJSON(responseJson.StatusCode, responseJson)
}

func (controller *UserControllerImpl) GetUsers(c *gin.Context) {
	data := controller.service.GetUsers(c.Request.Context())

	responseJson := response.WebResponse{
		StatusCode: http.StatusOK,
		Message:    "ok",
		Data:       data,
	}

	c.IndentedJSON(responseJson.StatusCode, responseJson)
}

func (controller *UserControllerImpl) Update(c *gin.Context) {

}
