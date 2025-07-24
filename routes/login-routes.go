package routes

import (
	"Tasktop/controllers"

	"github.com/gorilla/mux"
)

var LogRegister = func(r *mux.Router) {
	r.HandleFunc("/register", controllers.SignUp).Methods("POST")
	r.HandleFunc("/register", controllers.SignUp).Methods("GET")
	r.HandleFunc("/login", controllers.LogIn).Methods("POST")
	r.HandleFunc("/logout", controllers.LogOut).Methods("GET")
}
