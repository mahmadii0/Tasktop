package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func ConvertTimeToPlanId(time string) (int, error) {
	selected := time[0:10]
	for i, item := range selected {
		if item == '-' {
			t1 := selected[:i]
			t2 := selected[i+1:]
			selected = t1 + t2
		}
	}
	Id, err := strconv.Atoi(selected)
	return Id, err
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
