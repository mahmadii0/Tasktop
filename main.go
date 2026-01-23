package main

import (
	"Tasktop/configure"
	"Tasktop/controllers"
	"Tasktop/middlewares"
	"Tasktop/routes"
	"Tasktop/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	routes.MainRegister(router)
	routes.LogRegister(router)
	dashRouter := router.PathPrefix("/dashboard").Subrouter()
	routes.DashRegister(dashRouter)
	dashRouter.Use(middlewares.AuthMiddleware)
	configure.CreateTables()
	utils.Connect()
	go controllers.DNotes()
	log.Fatal(http.ListenAndServe("localhost:8080", router))

}
