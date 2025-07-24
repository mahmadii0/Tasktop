package routes

import (
	"Tasktop/controllers"

	"github.com/gorilla/mux"
)

var MainRegister = func(r *mux.Router) {
	r.HandleFunc("/", controllers.IndexHandler)
}
