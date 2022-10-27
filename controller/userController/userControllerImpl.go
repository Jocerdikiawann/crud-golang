package usercontroller

import (
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

func (controller *UserControllerImpl) AuthSignIn(ctx *gin.Context) {
	res := controller.service.AuthSignIn(ctx)
	
	ctx.IndentedJSON(res.StatusCode, res)
	return
}

func (controller *UserControllerImpl) Create(c *gin.Context) {
	data := controller.service.Create(c)

	c.IndentedJSON(data.StatusCode, data)
	return
}

func (controller *UserControllerImpl) GetUser(c *gin.Context) {
	data := controller.service.GetUser(c)

	c.IndentedJSON(data.StatusCode, data)
	return
}

func (controller *UserControllerImpl) GetUsers(c *gin.Context) {
	data := controller.service.GetUsers(c)

	c.IndentedJSON(data.StatusCode, data)
	return
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
		data, err := controller.service.Update(c.Request.Context(), body, id)
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
