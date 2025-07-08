package controllers

import (
	"Tasktop/middlewares"
	"Tasktop/models"
	"Tasktop/utils"
	"fmt"
	"net/http"
	"time"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user = &models.User{}
	var questions = &models.SecurityQuestions{}

	if r.Method == http.MethodGet {
		TemplateRender(w, "/main/authentication", nil)
	} else if r.Method == http.MethodPost {
		userName := r.FormValue("username")
		fullName := r.FormValue("fisrtName") + "  " + r.FormValue("lastName")
		email := r.FormValue("email")
		phone := r.FormValue("phone")
		password := r.FormValue("password")
		questions1 := r.FormValue("question1")
		answer1 := r.FormValue("answer1")
		questions2 := r.FormValue("question2")
		answer2 := r.FormValue("answer2")

		//Error Handling
		if len(password) < 8 || len(userName) < 8 {
			err := http.StatusNotAcceptable
			http.Error(w, "Invalid Password or Username", err)
		}
		u, _ := models.GetUserByUserName(userName)
		if u.FullName != "" {
			if u.UserName == userName || u.Email == email || u.Phone == phone {
				err := http.StatusConflict
				http.Error(w, "The user information(userName,Email,Phone) already used", err)
				return
			}
		}

		//Hashing
		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			er := http.StatusNotAcceptable
			http.Error(w, "Error While Hashing", er)
			return
		}
		hashedAnswer1, err := utils.HashPassword(answer1)
		if err != nil {
			er := http.StatusNotAcceptable
			http.Error(w, "Error While Hashing", er)
			return
		}
		hashedAnswer2, err := utils.HashPassword(answer2)
		if err != nil {
			er := http.StatusNotAcceptable
			http.Error(w, "Error While Hashing", er)
			return
		}

		//User Injection
		user.UserName = userName
		user.FullName = fullName
		user.Email = email
		user.Phone = phone
		user.Password = hashedPassword
		status := models.AddUser(user)
		if !(status) {
			err := http.StatusBadRequest
			http.Error(w, "Bad Request", err)
			return
		}

		//Questions Injection
		questions.Question1 = questions1
		questions.Answer1 = hashedAnswer1
		questions.Question2 = questions2
		questions.Answer2 = hashedAnswer2
		status = models.AddQuestions(user.UserName, questions)
		if !(status) {
			err := http.StatusBadRequest
			http.Error(w, "Bad Request", err)
			return
		}
		return

	} else {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid Method", err)
		return
	}

}

func LogIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid Method", err)
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		err := http.StatusBadRequest
		http.Error(w, "Invalid Request", err)
		return
	}

	//Check Email
	hash, err := models.GetPassHashByEmail(email)
	if err != nil {
		er := http.StatusBadRequest
		fmt.Printf("Wrong Email or Password")
		http.Error(w, "Error For Getting Hash password", er)
		return
	}
	status := utils.CheckPassword(password, hash)
	if !(status) {
		er := http.StatusBadRequest
		fmt.Printf("Wrong Email or Password")
		http.Error(w, "Wrong Password", er)
		return
	}

	sessionToken := utils.GenerateToken(32)
	csrfToken := utils.GenerateToken(32)
	//Set session cookie

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})

	status = models.SetTokens(sessionToken, csrfToken, email)
	if !(status) {
		er := http.StatusBadRequest
		http.Error(w, "Error while Set tokens", er)
		return
	}
	return

}
func LogOut(w http.ResponseWriter, r *http.Request) {
	if err := middlewares.Authorize(r); err != nil {
		er := http.StatusUnauthorized
		http.Error(w, "UnAuthorized", er)
		return
	}

	//Clear Cookies
	st, _ := r.Cookie("session_token")
	email := models.GetEmailBySessionToken(st.Value)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: false,
	})

	status := models.ClearTokens(email)
	if !(status) {
		er := http.StatusBadRequest
		http.Error(w, "Error while clear tokens", er)
		return
	}
	return
}
