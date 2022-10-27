package categorycontroller

// import (
// 	categorydomain "belajar-golang-rest-api/models/domain/categoryDomain"
// 	"belajar-golang-rest-api/models/response"
// 	categoryservices "belajar-golang-rest-api/services/categoryServices"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type CategoryControllerImpl struct {
// 	service categoryservices.CategoryService
// }

// func NewCategoryController(service categoryservices.CategoryService) CategoryController {
// 	return &CategoryControllerImpl{
// 		service: service,
// 	}
// }

// func (c *CategoryControllerImpl) Create(g *gin.Context) {
// 	//init body
// 	var body categorydomain.CategoryRequest

// 	if e := g.BindJSON(&body); e != nil {
// 		g.IndentedJSON(http.StatusBadRequest, response.WebResponse{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    e.Error(),
// 			Data:       gin.H{},
// 		})
// 		return
// 	}
// 	data, err := c.service.Create(g.Request.Context(), body)

// 	if err != nil {
// 		g.IndentedJSON(http.StatusBadRequest, response.WebResponse{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    err.Error(),
// 			Data:       gin.H{},
// 		})
// 		return
// 	}

// 	g.IndentedJSON(http.StatusCreated, response.WebResponse{
// 		StatusCode: http.StatusCreated,
// 		Message:    "Ok",
// 		Data:       data,
// 	})
// 	return
// }

// func (c *CategoryControllerImpl) GetCategory(g *gin.Context) {
// 	id := g.Param("id")
// 	data, err := c.service.GetCategory(g.Request.Context(), id)

// 	if err != nil {
// 		g.IndentedJSON(http.StatusNotFound, response.WebResponse{
// 			StatusCode: http.StatusNotFound,
// 			Message:    err.Error(),
// 			Data:       gin.H{},
// 		})
// 		return
// 	}

// 	g.IndentedJSON(http.StatusOK, response.WebResponse{
// 		StatusCode: http.StatusOK,
// 		Message:    "Ok",
// 		Data:       data,
// 	})
// 	return
// }

// func (c *CategoryControllerImpl) GetCategories(g *gin.Context) {
// 	data, err := c.service.GetCategories(g.Request.Context())

// 	if err != nil {
// 		g.IndentedJSON(http.StatusNotFound, response.WebResponse{
// 			StatusCode: http.StatusNotFound,
// 			Message:    err.Error(),
// 			Data:       make([]string, 0),
// 		})
// 		return
// 	}

// 	g.IndentedJSON(http.StatusOK, response.WebResponse{
// 		StatusCode: http.StatusOK,
// 		Message:    "Ok",
// 		Data:       data,
// 	})
// 	return
// }

// func (c *CategoryControllerImpl) Update(g *gin.Context) {

// }

// func (c *CategoryControllerImpl) Delete(g *gin.Context) {

// }
