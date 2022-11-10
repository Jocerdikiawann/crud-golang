package roleservices

// import (
// 	"belajar-golang-rest-api/models/response"
// 	"belajar-golang-rest-api/models/roles"
// 	rolerepositories "belajar-golang-rest-api/repository/roleRepositories"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/validator/v10"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type RoleServiceImpl struct {
// 	RoleRepo rolerepositories.RoleRepository
// 	Db       *mongo.Database
// 	Validate *validator.Validate
// }

// func NewRoleService(repo rolerepositories.RoleRepository, db *mongo.Database, validator *validator.Validate) RoleService {
// 	return &RoleServiceImpl{
// 		RoleRepo: repo,
// 		Db:       db,
// 		Validate: validator,
// 	}
// }

// func (s *RoleServiceImpl) Create(ctx *gin.Context) response.Response {
// 	var req roles.RolesReq

// 	bindJson := ctx.BindJSON(&req)

// 	if bindJson != nil {
// 		return response.Response{
// 			Code:    http.StatusBadRequest,
// 			Message: bindJson.Error(),
// 			Data:    gin.H{},
// 		}
// 	}

// 	validate := s.Validate.Struct(&req)

// 	if validate != nil {
// 		return response.Response{
// 			Code:    http.StatusBadRequest,
// 			Message: validate.Error(),
// 			Data:    gin.H{},
// 		}
// 	}

// 	res, err := s.RoleRepo.Create(ctx, s.Db, req)

// 	if err != nil {
// 		return response.Response{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 			Data:    gin.H{},
// 		}
// 	}

// 	return response.Response{
// 		Code:    http.StatusCreated,
// 		Message: "created",
// 		Data:    res,
// 	}
// }
