package middlewares

import (
	"encoding/json"
	"github.com/Achmadqizwini/SportKai/config"
	"github.com/Achmadqizwini/SportKai/utils/helper"
	"net/http"
	"strings"
	"time"

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

// Middleware function for JWT authentication
func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the JWT token from the Authorization header
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		if tokenString == "" {
			json.NewEncoder(w).Encode(helper.FailedResponse("You are not authorized for this operations. Login first"))

			return
		}
		tokenString = tokenString[len("Bearer "):]

		// Parse the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method and return the secret key
			return secretKey, nil
		})
		if err != nil || !token.Valid {
			json.NewEncoder(w).Encode(helper.FailedResponse("You are not authorized for this operations. Login first"))
			return
		}
		// Extract claims from the token
		// claims, ok := token.Claims.(jwt.MapClaims)
		// if !ok {
		// 	json.NewEncoder(w).Encode(helper.FailedResponse("You are not authorized for this operations"))
		// 	return
		// }

		// Access specific data from the claims
		// userI := claims["public_id"].(string)
		// username := claims["username"].(string)
		// email := claims["email"].(string)

		// // Optionally, pass the extracted data to the request context
		// r = r.WithContext(context.WithValue(r.Context(), "user_id", userID))
		// r = r.WithContext(context.WithValue(r.Context(), "username", username))
		// r = r.WithContext(context.WithValue(r.Context(), "email", email))

		// Call the next handler in the chain
		next(w, r)
	}
}
