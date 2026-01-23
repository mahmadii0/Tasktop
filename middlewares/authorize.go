package middlewares

import (
	"Tasktop/models"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var AuthError = errors.New("Unauthorized")

func Authorize(r *http.Request) error {
	tokenStr, err := r.Cookie("auth")
	if err != nil || tokenStr.Value == "" {
		return AuthError
	}

	jwtSecret := os.Getenv("SECRETJWT")
	if jwtSecret == "" {
		fmt.Errorf("JWT secret not configured")
		return AuthError
	}

	token, err := jwt.Parse(tokenStr.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return AuthError
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return AuthError
	}

	ttl, _ := claims["ttl"].(float64)

	if int64(ttl) < time.Now().Unix() {
		return AuthError
	}

	if jwtSecret == "" {
		return fmt.Errorf("JWT secret not configured")
	}

	userIdFloat, ok := claims["userId"].(float64)
	if !ok {
		return AuthError
	}
	userId := int64(userIdFloat)
	user := models.UserFromId(userId)
	if user.ID == 0 {
		return AuthError
	}

	return nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := Authorize(r)
		if err != nil {
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
