package models

import (
	"Tasktop/configure"
	"Tasktop/utils"
	"database/sql"
)

type User struct {
	UserID   int
	Name     string
	Email    string
	Phone    string
	Password string
}

var (
	db *sql.DB
)

func init() {
	configure.Connect()
	db = configure.GetDB()
}

func GetUserById(userId int) (*User, error) {
	var user User
	query := `SELECT userId,name,email,phone FROM users WHERE userId=?`
	row := db.QueryRow(query, userId)
	err := row.Scan(&user.UserID, &user.Name, &user.Email, &user.Phone)
	return &user, err
}

func (user *User) AddUser() (*User, bool) {
	status := true
	query := `INSERT INTO users(name,email,password) VALUES (?,?,?)`
	pass, err := utils.HashPassword(user.Password)
	if err != nil {
		status = false
	}
	result, err := db.Exec(query, user.Name, user.Email, pass)
	if err != nil {
		status = false
	}
	id, err := result.LastInsertId()
	if err != nil {
		status = false
	}
	user.UserID = int(id)
	return user, status
}

func UpdateUser(user *User) bool {
	status := true
	query := `UPDATE users SET name= ?, email= ?, phone = ? WHERE id = ?`
	_, err := db.Exec(query, user.Name, user.Email, user.Phone, user.UserID)
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
