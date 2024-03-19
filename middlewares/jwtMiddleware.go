package middlewares

import (
	"time"

	"github.com/Achmadqizwini/SportKai/config"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey []byte

func InitJWT(c *config.AppConfig) {
	secretKey = []byte(c.AppSecretKey)
}

func CreateToken(id uint, publicID string, username string, email string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":        id,
			"public_id": publicID,
			"username":  username,
			"email":     email,
			"exp":       time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
