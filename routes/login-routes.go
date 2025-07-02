package routes

import (
	"Tasktop/controllers"

	"github.com/gorilla/mux"
)

var LogRegister = func(r *mux.Router) {
	r.HandleFunc("/register", controllers.SignUp)
	r.HandleFunc("/login", controllers.LogIn)
	r.HandleFunc("/logout", controllers.LogOut)
}
