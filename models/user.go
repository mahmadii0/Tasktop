package models

import (
	"Tasktop/configure"
	"database/sql"
)

type User struct {
	UserName     string `json:"username"`
	FullName     string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
	SessionToken string
	CSRF         string
}

type SecurityQuestions struct {
	UserName  string `json:"username"`
	Question1 string `json:"question1"`
	Answer1   string `json:"answer1"`
	Question2 string `json:"question2"`
	Answer2   string `json:"answer2"`
}

var (
	db *sql.DB
)

func init() {
	configure.Connect()
	db = configure.GetDB()
}

//User

func AddUser(user *User) bool {
	status := true
	query := `INSERT INTO users(username,fullname,email,phone,password) VALUES(?,?,?,?,?)`
	_, err := db.Exec(query, user.UserName, user.FullName, user.Email, user.Phone, user.Password)
	if err != nil {
		status = false
	}
	return status
}

func SetTokens(sessionToken string, csrfToken string, email string) bool {
	status := true
	query := `UPDATE users SET session_token=? ,csrf_token=? WHERE email=?`
	_, err := db.Exec(query, sessionToken, csrfToken, email)
	if err != nil {
		status = false
	}
	return status
}

func GetEmailBySessionToken(sessionToken string) string {
	email := ""
	query := `SELECT email From users WHERE session_token=?`
	row := db.QueryRow(query, sessionToken)
	_ = row.Scan(&email)
	return email
}

func CompareCsrfToken(email string, csrf string) bool {
	status := true
	csrfDb := ""
	query := `SELECT csrf_token From users WHERE email=?`
	row := db.QueryRow(query, email)
	_ = row.Scan(&csrfDb)
	if csrfDb != csrf {
		status = false
	}
	return status
}

func GetPassHashByEmail(email string) (string, error) {
	var hash string
	query := `SELECT password FROM users WHERE email=?`
	row := db.QueryRow(query, email)
	err := row.Scan(&hash)
	return hash, err
}

func GetUserByUserName(username string) (*User, error) {
	var user User
	query := `SELECT username,fullname,email,phone FROM users WHERE username=?`
	row := db.QueryRow(query, username)
	err := row.Scan(&user.UserName, &user.FullName, &user.Email, &user.Phone)
	return &user, err
}

func UpdateUser(user *User) bool {
	status := true
	query := `UPDATE users SET name= ?, email= ?, phone = ? WHERE id = ?`
	_, err := db.Exec(query, user.FullName, user.Email, user.Phone, user.UserName)
	if err != nil {
		status = false
	}
	return status
}

func DeleteUser(username int) bool {
	status := true
	query := `DELETE FROM users WHERE id= ?`
	_, err := db.Exec(query, username)
	if err != nil {
		status = false
	}
	return status
}

//Questions

func AddQuestions(username string, questions *SecurityQuestions) bool {
	status := true
	query := `INSERT INTO securityquestions VALUES(?,?,?,?,?)`
	_, err := db.Exec(query, username, questions.Question1, questions.Answer1, questions.Question2, questions.Answer2)
	if err != nil {
		status = false
	}
	return status
}
