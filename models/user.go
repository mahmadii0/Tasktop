package models

import (
	"Tasktop/configure"
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName     string `gorm:"primaryKey;size:100" json:"username"`
	FullName     string `gorm:"not null;size:150" json:"name"`
	Email        string `gorm:"unique;not null;size250" json:"email"`
	Phone        string `gorm:"unique;not null;size:13" json:"phone"`
	Password     string `gorm:"not null,size:400" json:"password"`
	SessionToken string `gorm:"size:64"`
	CSRF         string `gorm:"size:64"`
}

type SecurityQuestions struct {
	gorm.Model
	UserName  string `gorm:"not null;size:100" json:"user_name"`
	Question1 string `gorm:"not null,size:100" json:"question1"`
	Answer1   string `gorm:"not null,size:2" json:"answer1"`
	Question2 string `gorm:"not null,size:100" json:"question2"`
	Answer2   string `gorm:"not null,size:2" json:"answer2"`
	User      User   `gorm:"foreignKey:UserName"`
}

var (
	db  *gorm.DB
	ctx context.Context
)

func init() {
	db, ctx = configure.GetDBctx()
}

//User

func AddUser(user *User) bool {
	gorm.G[User](db).Create(ctx, user)
	return true
}

func SetTokens(sessionToken string, csrfToken string, email string) bool {
	gorm.G[User](db).Where("email=?", email).Updates(ctx, User{CSRF: csrfToken, SessionToken: sessionToken})
	return true
}

func ClearTokens(email string) bool {
	gorm.G[User](db).Where("email=?", email).Updates(ctx, User{CSRF: "", SessionToken: ""})
	return true
}

//func GetEmailBySessionToken(sessionToken string) string {
//	email := ""
//	gorm.G[User](db).Where("")
//	query := `SELECT email From users WHERE session_token=?`
//	row := db.QueryRow(query, sessionToken)
//	_ = row.Scan(&email)
//	return email
//}

//
//func GetUsernameBySessionToken(sessionToken string) string {
//	email := ""
//	query := `SELECT username From users WHERE session_token=?`
//	row := db.QueryRow(query, sessionToken)
//	_ = row.Scan(&email)
//	return email
//}
//
//func CompareCsrfToken(email string, csrf string) bool {
//	status := true
//	csrfDb := ""
//	query := `SELECT csrf_token From users WHERE email=?`
//	row := db.QueryRow(query, email)
//	_ = row.Scan(&csrfDb)
//	if csrfDb != csrf {
//		status = false
//	}
//	return status
//}
//
//func GetPassHashByEmail(email string) (string, error) {
//	var hash string
//	query := `SELECT password FROM users WHERE email=?`
//	row := db.QueryRow(query, email)
//	err := row.Scan(&hash)
//	return hash, err
//}
//
//func GetUserByUserName(username string) (*User, error) {
//	var user User
//	query := `SELECT username,fullname,email,phone FROM users WHERE username=?`
//	row := db.QueryRow(query, username)
//	err := row.Scan(&user.UserName, &user.FullName, &user.Email, &user.Phone)
//	return &user, err
//}
//
//func UpdateUser(user *User) bool {
//	status := true
//	query := `UPDATE users SET name= ?, email= ?, phone = ? WHERE id = ?`
//	_, err := db.Exec(query, user.FullName, user.Email, user.Phone, user.UserName)
//	if err != nil {
//		status = false
//	}
//	return status
//}
//
//func DeleteUser(username int) bool {
//	status := true
//	query := `DELETE FROM users WHERE id= ?`
//	_, err := db.Exec(query, username)
//	if err != nil {
//		status = false
//	}
//	return status
//}
//
////Questions
//
//func AddQuestions(username string, questions *SecurityQuestions) bool {
//	status := true
//	query := `INSERT INTO securityquestions VALUES(?,?,?,?,?)`
//	_, err := db.Exec(query, username, questions.Question1, questions.Answer1, questions.Question2, questions.Answer2)
//	if err != nil {
//		status = false
//	}
//	return status
//}
