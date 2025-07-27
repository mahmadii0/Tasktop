package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func SeparateDateTime(dateTime string) (string, string) {
	date := dateTime[0:10]
	time := dateTime[12:]
	return date, time
}

func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	return string(bytes), err
}

func CheckPassword(pass string, hash string) bool {
	status := true
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		status = false
	}
	return status
}

func GenerateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Error while creating Token")
	}
	return base64.URLEncoding.Strict().EncodeToString(bytes)
}
