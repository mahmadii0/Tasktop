package main

import (
	"Tasktop/configure"
	"Tasktop/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	routes.MainRegister(router)
	routes.DashRegister(router)
	routes.LogRegister(router)
	configure.CreateTables()
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
