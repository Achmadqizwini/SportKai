package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Achmadqizwini/SportKai/config"
	"github.com/Achmadqizwini/SportKai/utils/helper"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey []byte

func InitJWT(c *config.AppConfig) {
	secretKey = []byte(c.AppSecretKey)
}

type Val string

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

		if err != nil {
			if err == jwt.ErrTokenExpired {
				json.NewEncoder(w).Encode(helper.FailedResponse("Access Denied: Your current session has expired. Please log in again to continue."))
				return
			} else if err == jwt.ErrSignatureInvalid {
				json.NewEncoder(w).Encode(helper.FailedResponse("Access Denied: Invalid token signature"))
				return
			} else {
				json.NewEncoder(w).Encode(helper.FailedResponse("Error parsing token: " + err.Error()))
				return
			}
		}

		if !token.Valid {
			json.NewEncoder(w).Encode(helper.FailedResponse("Access Denied: Token is invalid"))
			return
		}

		// Extract claims from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			json.NewEncoder(w).Encode(helper.FailedResponse("You are not authorized for this operations. Login first"))
			return
		}

		// // Optionally, pass the extracted data to the request context
		r = r.WithContext(context.WithValue(r.Context(), Val("user_id"), claims["public_id"].(string)))
		r = r.WithContext(context.WithValue(r.Context(), Val("username"), claims["username"].(string)))
		r = r.WithContext(context.WithValue(r.Context(), Val("email"), claims["email"].(string)))

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	}
}
