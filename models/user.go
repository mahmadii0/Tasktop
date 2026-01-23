package models

import (
	"Tasktop/utils"
	"context"

	"gorm.io/gorm"
)

type User struct {
	ID           int64  `gorm:"primaryKey;autoIncrement" json:"userId"`
	UserName     string `gorm:"size:100;" json:"username"`
	FullName     string `gorm:"not null;size:150" json:"name"`
	Email        string `gorm:"unique;not null;size:250" json:"email"`
	Phone        string `gorm:"unique;not null;size:13" json:"phone"`
	Password     string `gorm:"not null;size:400" json:"password"`
	SessionToken string `gorm:"size:64"`
	CSRF         string `gorm:"size:64"`
}

type SecurityQuestions struct {
	UserID    int64  `gorm:"not null" json:"userId"`
	Question1 string `gorm:"not null;size:100" json:"question1"`
	Answer1   string `gorm:"not null;size:100" json:"answer1"`
	Question2 string `gorm:"not null;size:100" json:"question2"`
	Answer2   string `gorm:"not null;size:100" json:"answer2"`
	User      User   `gorm:"foreignKey:UserID"`
}

var (
	db  *gorm.DB
	ctx context.Context
)

func init() {
	db, ctx = utils.GetDBctx()
}

//User

func AddUser(user *User) bool {
	result := gorm.G[User](db).Create(ctx, user)
	if result != nil {
		return false
	}
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
	u, err := gorm.G[User](db).Where("session_token=?", sessionToken).Select("email").Find(ctx)
	if err != nil {
		return ""
	}
	email = u[0].Email
	return email
}

func GetUserIdBySessionToken(sessionToken string) int64 {
	var userId int64
	u, err := gorm.G[User](db).Where("session_token=?", sessionToken).Select("user_name").Find(ctx)
	if err != nil {
		return 0
	}
	userId = u[0].ID
	return userId
}

func CompareCsrfToken(email string, csrf string) bool {
	u, err := gorm.G[User](db).Where("email=?", email).Select("csrf").Find(ctx)
	if err != nil {
		return false
	}
	csrfDb := ""
	csrfDb = u[0].CSRF
	if csrfDb != csrf {
		return false
	}
	return true
}

func GetPassHashByEmail(email string) (string, error) {
	var hash string
	u, err := gorm.G[User](db).Where("email=?", email).Select("password").Find(ctx)
	if err != nil {
		return "", err
	}
	hash = u[0].Password
	return hash, err
}

func GetUserByUserName(username string) (*User, error) {
	var user User
	u, err := gorm.G[User](db).Where("user_name=?", username).Select("user_name,full_name,email,phone").Find(ctx)
	if err != nil {
		return nil, err
	}
	if len(u) == 0 {
		return &user, err
	}
	user = u[0]
	return &user, err
}

func UpdateUser(user *User) bool {
	gorm.G[User](db).Where("user_name=?", user.UserName).Updates(ctx, User{FullName: user.FullName, Email: user.Email, Phone: user.Phone})
	return true
}

func DeleteUser(username int) bool {
	gorm.G[User](db).Where("user_name = ?", username).Delete(ctx)
	return true
}

//Questions

func AddQuestions(questions *SecurityQuestions) bool {
	result := gorm.G[SecurityQuestions](db).Create(ctx, questions)
	if result != nil {
		return false
	}
	return true
}
