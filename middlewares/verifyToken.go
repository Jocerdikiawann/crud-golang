// Package middlewares provides ...
package middlewares

import (
	"belajar-golang-rest-api/models/response"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func MiddlewareAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		res := response.WebResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "Unauthorized",
			Data:       gin.H{},
		}

		header := c.GetHeader("Authorization")
		bearerToken := strings.Split(header, " ")

		if len(bearerToken) == 2 {

			token, _ := jwt.Parse(header, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
				}
				return []byte(os.Getenv("SECRET_TOKEN")), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				c.Set("_id", claims["_id"])
			} else {

				c.JSON(res.StatusCode, res)
				c.Abort()
				return
			}
		} else {
			c.JSON(res.StatusCode, res)
			c.Abort()
			return
		}
	}
}
