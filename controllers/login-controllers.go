package controllers

import (
	"Tasktop/middlewares"
	"Tasktop/models"
	"Tasktop/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user = &models.User{}
	var questions = &models.SecurityQuestions{}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if r.Method == http.MethodGet {
		htmlContent, err := os.ReadFile(os.Getenv("TEMPLATES_SOURCE") + "/main/authentication.html")
		if err != nil {
			http.Error(w, "Error reading HTML file", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(htmlContent)
	} else if r.Method == http.MethodPost {
		userName := r.FormValue("username")
		userName = strings.ToLower(userName)
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
			return
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
		questions.UserID = user.ID
		status = models.AddQuestions(questions)
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
	// if r.Method != http.MethodPost {
	// 	err := http.StatusMethodNotAllowed
	// 	http.Error(w, "Invalid Method", err)
	// 	return
	// }
	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		err := http.StatusBadRequest
		http.Error(w, "Invalid Request", err)
		return
	}

	//Check Email
	userId, hash, err := models.GetPassHashByEmail(email)
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

	tokenString := utils.GenerateJWT(email, userId)
	if tokenString == "" {
		fmt.Printf("Failed to Generate JWT")
		http.Error(w, "Failed to Generate JWT", http.StatusBadGateway)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24 * 1),
		HttpOnly: true,
	})

}
func LogOut(w http.ResponseWriter, r *http.Request) {
	if err := middlewares.Authorize(r); err != nil {
		er := http.StatusUnauthorized
		http.Error(w, "UnAuthorized", er)
		return
	}

	//Clear Cookies

	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	return
}
