package middlewares

import (
	"Tasktop/models"
	"errors"
	"net/http"
)

var AuthError = errors.New("Unauthorized")

func Authorize(r *http.Request) error {
	st, err := r.Cookie("session_token")
	if err != nil || st.Value == "" {
		return AuthError
	}
	email := models.GetEmailBySessionToken(st.Value)

	csrf, err := r.Cookie("csrf_token")
	if csrf.Value == "" || err != nil {
		return AuthError
	}
	status := models.CompareCsrfToken(email, csrf.Value)
	if !(status) {
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
