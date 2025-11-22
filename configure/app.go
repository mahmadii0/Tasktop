package configure

import (
	"Tasktop/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
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
