package utils

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Layouts = map[string]string{
	"RFC3339":   time.RFC3339, // "2006-01-02T15:04:05Z07:00"
	"RFC1123":   time.RFC1123, // "Mon, 02 Jan 2006 15:04:05 MST"
	"Date only": "2006-01-02", // "2023-12-25"
	"DateTime":  "2006-01-02 15:04:05",
	"US format": "01/02/2006", // "12/25/2023"
	"Time only": "15:04:05",
}

var (
	db  *gorm.DB
	ctx context.Context
)

func BoolToInt(b bool) int {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}

func SeparateDateTime(dateTime string) (string, string) {
	date := dateTime[0:11]
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

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	source := os.Getenv("DATABASE_SOURCE")
	d, err := gorm.Open(mysql.Open(source), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	ctx = context.Background()
	db = d
}

func GetDBctx() (*gorm.DB, context.Context) {
	Connect()
	return db, ctx
}

func ParseTime(layout string, t string) (time.Time, error) {
	layout = Layouts[layout]
	time, err := time.Parse(layout, t)
	if err != nil {
		return time, err
	}
	return time, nil
}
