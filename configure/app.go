package configure

import (
	"Tasktop/models"
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	db  *gorm.DB
	ctx context.Context
)

func CreateTables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	source := os.Getenv("DATABASE_SOURCE")
	db, err := gorm.Open(mysql.Open(source), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.DailyPlan{}, models.DailyGoal{},
		models.MonthlyGoal{}, models.MonthlyPlan{}, models.AnnuallyGoal{},
		models.AnnuallyPlan{}, models.Note{}, models.User{}, models.SecurityQuestions{})
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
	return db, ctx
}
