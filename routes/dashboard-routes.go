package routes

import (
	"Tasktop/controllers"
	"Tasktop/models"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

var AuthError = errors.New("Unauthorized")

func Authorize(r *http.Request) error {
	st, err := r.Cookie("session_token")
	if err != nil || st.Value == "" {
		return AuthError
	}
	email := models.GetEmailBySessionToken(st.Value)

	csrf := r.Header.Get("X-CSRF-Token")
	if csrf == "" {
		return AuthError
	}
	status := models.CompareCsrfToken(email, csrf)
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

var DashRegister = func(r *mux.Router) {
	r.HandleFunc("/", controllers.DashHandler)
	//Read Goals
	r.HandleFunc("/daily-goals/{goalId:[0-99]+}", controllers.DailyGoal).Methods("GET")
	r.HandleFunc("/monthly-goals/{goalId:[0-99]+}", controllers.MonthlyGoal).Methods("GET")
	r.HandleFunc("/annually-goals/{goalId:[0-99]+}", controllers.AnnuallyGoal).Methods("GET")

	//Create Goals
	r.HandleFunc("/daily-goals/create", controllers.CDailyGoal).Methods("POST")
	r.HandleFunc("/monthly-goals/create", controllers.CMonthlyGoal).Methods("POST")
	r.HandleFunc("/annually-goals/create", controllers.CAnnuallyGoal).Methods("POST")

	//Update Goals
	r.HandleFunc("/daily-goals/{goalId:[0-99]+}", controllers.UDailyGoal).Methods("PUT")
	r.HandleFunc("/monthly-goals/{goalId:[0-99]+}", controllers.UMonthlyGoal).Methods("PUT")
	r.HandleFunc("/annually-goals/{goalId:[0-99]+}", controllers.UAnnuallyGoals).Methods("PUT")

	//Delete Goals
	r.HandleFunc("/daily-goals/{goalId:[0-99]+}", controllers.DDailyGoal).Methods("DELETE")
	r.HandleFunc("/monthly-goals/{goalId:[0-99]+}", controllers.DMonthlyGoal).Methods("DELETE")
	r.HandleFunc("/annually-goals/{goalId:[0-99]+}", controllers.DAnnuallyGoal).Methods("DELETE")

	//Report
	r.HandleFunc("/report", controllers.Report)
}
