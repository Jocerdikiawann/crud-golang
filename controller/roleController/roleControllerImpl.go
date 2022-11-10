package rolecontroller

// import (
// 	roleservices "belajar-golang-rest-api/services/roleServices"

// 	"github.com/gin-gonic/gin"
// )

// type RoleControllerImpl struct {
// 	service roleservices.RoleService
// }

// func NewRoleController(service roleservices.RoleService) RoleController {
// 	return &RoleControllerImpl{
// 		service: service,
// 	}
// }

// func (c *RoleControllerImpl) Create(ctx *gin.Context) {
// 	res := c.service.Create(ctx)

// 	ctx.IndentedJSON(res.Code, res)
// 	return
// }
