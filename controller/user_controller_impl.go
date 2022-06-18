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
	var responseJson response.WebResponse
	if e := c.BindJSON(&body); e != nil {
		responseJson = response.WebResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "error",
			Data:       gin.H{},
		}
	} else {
		data, err := controller.service.Create(c.Request.Context(), body)
		if err != nil {
			responseJson = response.WebResponse{
				StatusCode: http.StatusNotFound,
				Message:    "error",
				Data:       gin.H{},
			}
		} else {
			responseJson = response.WebResponse{
				StatusCode: http.StatusCreated,
				Message:    "ok",
				Data:       data,
			}
		}
	}

	c.IndentedJSON(http.StatusCreated, responseJson)
}

func (controller *UserControllerImpl) GetUser(c *gin.Context) {
	var responseJson response.WebResponse
	id := c.Param("id")
	data, err := controller.service.GetUser(c.Request.Context(), id)

	if err != nil {
		responseJson = response.WebResponse{
			StatusCode: http.StatusNotFound,
			Message:    "error",
			Data:       gin.H{},
		}
	} else {
		responseJson = response.WebResponse{
			StatusCode: http.StatusOK,
			Message:    "ok",
			Data:       data,
		}
	}

	c.IndentedJSON(responseJson.StatusCode, responseJson)
}

func (controller *UserControllerImpl) GetUsers(c *gin.Context) {
	var responseJson response.WebResponse
	data, err := controller.service.GetUsers(c.Request.Context())

	if err != nil {
		responseJson = response.WebResponse{
			StatusCode: http.StatusNotFound,
			Message:    "error",
			Data:       gin.H{},
		}
	} else {
		responseJson = response.WebResponse{
			StatusCode: http.StatusOK,
			Message:    "ok",
			Data:       data,
		}
	}

	c.IndentedJSON(responseJson.StatusCode, responseJson)
}

func (controller *UserControllerImpl) Update(c *gin.Context) {

}
