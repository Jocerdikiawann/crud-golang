// Package middlewares provides ...
package middlewares

import (
	"belajar-golang-rest-api/models/response"
	"belajar-golang-rest-api/utils"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func MiddlewareAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		token, err := jwt.Parse(header, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_TOKEN")), nil
		})
		utils.IfErrorHandler(err)
		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			res := response.WebResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized",
				Data:       gin.H{},
			}
			c.JSON(res.StatusCode, res)
			c.Abort()
		}
		c.Next()
	}
}
