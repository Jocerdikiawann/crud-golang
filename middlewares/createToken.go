package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(_id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["_id"] = _id
	atClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, errSignToken := at.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
	if errSignToken != nil {
		return "", errSignToken
	}
	return accessToken, nil
}
