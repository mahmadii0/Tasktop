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

func GetEmailBySessionToken(sessionToken string) string {
	var email string
	e, err := gorm.G[User](db).Where("session_token=?", sessionToken).Select("email").Find(ctx)
	if err != nil {
		return ""
	}
	email = e[0].Email
	return email
}

func GetUsernameBySessionToken(sessionToken string) string {
	var username string
	u, err := gorm.G[User](db).Where("session_token=?", sessionToken).Select("username").Find(ctx)
	if err != nil {
		return ""
	}
	username = u[0].UserName
	return username
}

func CompareCsrfToken(email string, csrf string) bool {
	d, err := gorm.G[User](db).Where("email=?", email).Select("csrf_token").Find(ctx)
	if err != nil {
		return false
	}
	csrfDb := ""
	csrfDb = d[0].CSRF
	if csrfDb != csrf {
		return false
	}
	return true
}

func GetPassHashByEmail(email string) (string, error) {
	var hash string
	pass, err := gorm.G[User](db).Where("email=?", email).Select("password").Find(ctx)
	if err != nil {
		return "", err
	}
	hash = pass[0].Password
	return hash, err
}

func GetUserByUserName(username string) (*User, error) {
	var user User
	u, err := gorm.G[User](db).Where("username=?", username).Select("username,fullname,email,phone").Find(ctx)
	if err != nil {
		return nil, err
	}
	user = u[0]
	return &user, err
}

func UpdateUser(user *User) bool {
	gorm.G[User](db).Where("username=?", user.UserName).Updates(ctx, User{FullName: user.FullName, Email: user.Email, Phone: user.Phone})
	return true
}

func DeleteUser(username int) bool {
	gorm.G[User](db).Where("username = ?", username).Delete(ctx)
	return true
}

//Questions

func AddQuestions(username string, questions *SecurityQuestions) bool {
	gorm.G[SecurityQuestions](db).Create(ctx, questions)
	return true
}
