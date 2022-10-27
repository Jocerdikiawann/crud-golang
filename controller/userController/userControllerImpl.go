package usercontroller

import (
	userservices "belajar-golang-rest-api/services/userServices"

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
	data := controller.service.Update(c)
	c.IndentedJSON(data.StatusCode, data)
	return
}

func (controller *UserControllerImpl) Delete(c *gin.Context) {
	data := controller.service.Delete(c)
	c.IndentedJSON(data.StatusCode, data)
	return
}
