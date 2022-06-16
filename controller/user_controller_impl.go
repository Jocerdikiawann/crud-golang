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
	newUser := domain.User{
		Id:        data.Id,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Address:   data.Address,
	}
	responseJson := response.WebResponse{
		StatusCode: http.StatusCreated,
		Message:    "ok",
		Data:       newUser,
	}

	c.IndentedJSON(http.StatusCreated, responseJson)

}

func (controller *UserControllerImpl) Getuser(c *gin.Context) {

}
