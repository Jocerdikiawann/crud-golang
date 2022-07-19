package usercontroller

import (
	usersdomain "belajar-golang-rest-api/models/domain/usersDomain"
	userrequests "belajar-golang-rest-api/models/requests/userRequests"
	"belajar-golang-rest-api/models/response"
	userservices "belajar-golang-rest-api/services/userServices"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	service userservices.UserService
}

func NewUserController(services userservices.UserService) UserController {
	return &UserControllerImpl{
		services,
	}
}

func (controller *UserControllerImpl) Create(c *gin.Context) {
	var body usersdomain.User
	var responseJson response.WebResponse
	if e := c.BindJSON(&body); e != nil {
		responseJson = response.WebResponse{
			StatusCode: http.StatusBadRequest,
			Message:    e.Error(),
			Data:       gin.H{},
		}
	} else {
		data, err := controller.service.Create(c.Request.Context(), body)
		if err != nil {
			responseJson = response.WebResponse{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Data:       gin.H{},
			}
		} else {
			responseJson = response.WebResponse{
				StatusCode: http.StatusCreated,
				Message:    "Ok",
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
			Message:    err.Error(),
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
			Message:    err.Error(),
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

	id := c.Param("id")
	var body userrequests.UserBody
	var res response.WebResponse

	if e := c.BindJSON(&body); e != nil {
		res = response.WebResponse{
			StatusCode: http.StatusBadRequest,
			Message:    e.Error(),
			Data:       gin.H{},
		}
	} else {
		req := usersdomain.User{
			Id:        id,
			Email:     body.Email,
			Password:  body.Password,
			FirstName: body.FirstName,
			LastName:  body.LastName,
			Address:   body.Address,
		}
		data, err := controller.service.Update(c.Request.Context(), req)
		if err != nil {
			res = response.WebResponse{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
				Data:       gin.H{},
			}

		} else {
			res = response.WebResponse{
				StatusCode: http.StatusCreated,
				Message:    "Ok",
				Data:       data,
			}
		}
	}

	c.IndentedJSON(res.StatusCode, res)
}

func (controller *UserControllerImpl) Delete(c *gin.Context) {
	var res response.WebResponse
	id := c.Param("id")

	data, err := controller.service.Delete(c.Request.Context(), id)
	if err != nil {
		res = response.WebResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       gin.H{},
		}
	} else {
		res = response.WebResponse{
			StatusCode: http.StatusOK,
			Message:    "ok",
			Data:       data,
		}
		c.IndentedJSON(res.StatusCode, res)
	}
}
