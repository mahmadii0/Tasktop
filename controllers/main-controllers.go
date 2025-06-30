package controllers

import (
	"Tasktop/constants"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func TemplateRender(w http.ResponseWriter, templ string, data interface{}) {

	templates := []string{
		filepath.Join(constants.TEMPLATES_SOURCE + templ + ".html"),
		filepath.Join(constants.TEMPLATES_SOURCE + "/base.html"),
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

	TemplateRender(w, "/main/index", nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {

}
