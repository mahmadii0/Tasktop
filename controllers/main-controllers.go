package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func TemplateRender(w http.ResponseWriter, templ string, data interface{}) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	templates := []string{
		filepath.Join(os.Getenv("TEMPLATES_SOURCE") + templ + ".html"),
		filepath.Join(os.Getenv("TEMPLATES_SOURCE") + "/base.html"),
	}

	t, err := template.ParseFiles(templates...)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while passing tamplates : %v", err), http.StatusInternalServerError)
	}
	err = t.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while Executing template : %v", err), http.StatusInternalServerError)
	}

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/main/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {

}
