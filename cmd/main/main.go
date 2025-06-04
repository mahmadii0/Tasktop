package main

import (
	"Tasktop/configure"
	"Tasktop/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	routes.MainRegister(router)
	routes.DashRegister(router)
	configure.CreateTables()
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
