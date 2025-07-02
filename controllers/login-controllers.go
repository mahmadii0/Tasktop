package controllers

import (
	"Tasktop/models"
	"Tasktop/utils"
	"encoding/json"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	if r.Method == http.MethodGet {
		TemplateRender(w, "/main/sign-in", nil)
	} else if r.Method == http.MethodPost {
		userName := r.FormValue("username")
		fullName := r.FormValue("fullname")
		email := r.FormValue("email")
		phone := r.FormValue("phone")
		password := r.FormValue("password")
		//Error Handling
		if len(password) < 8 || len(userName) < 8 {
			err := http.StatusNotAcceptable
			http.Error(w, "Invalid Password or Username", err)
		}
		u, _ := models.GetUserByUserName(userName)
		if u.UserName == userName || u.Email == email || u.Phone == phone {
			err := http.StatusConflict
			http.Error(w, "The user information(userName,Email,Phone) already used", err)
		}
		//Hashing
		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			er := http.StatusNotAcceptable
			http.Error(w, "Error While Hashing", er)
		}

		user.UserName = userName
		user.FullName = fullName
		user.Email = email
		user.Phone = phone
		user.Password = hashedPassword
		user, status := models.AddUser(user)
		if !(status) {
			err := http.StatusBadRequest
			http.Error(w, "Bad Request", err)
		}
		json.NewEncoder(w).Encode(user)
		return //Create a welcome register page

	} else {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid Method", err)
	}

}

func LogIn(w http.ResponseWriter, r *http.Request) {

}
func LogOut(w http.ResponseWriter, r *http.Request) {

}
