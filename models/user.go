package models

import (
	"Tasktop/configure"
	"database/sql"
)

type User struct {
	UserName string `json:"username"`
	FullName string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
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

func AddUser(user *User) (*User, bool) {
	user.Email = "dfdfd"
	return user, true
}

func GetUserByUserName(username string) (*User, error) {
	var user User
	query := `SELECT userId,name,email,phone FROM users WHERE userId=?`
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

func DeleteUser(userId int) bool {
	status := true
	query := `DELETE FROM users WHERE id= ?`
	_, err := db.Exec(query, userId)
	if err != nil {
		status = false
	}
	return status
}

//Questions

func AddQuestions(user *User) bool {
	user.Email = "dfdfd"
	return true
}
