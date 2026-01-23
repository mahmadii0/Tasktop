package configure

import (
	"Tasktop/models"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateTables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	source := os.Getenv("DATABASE_SOURCE")
	db, err := gorm.Open(mysql.Open(source), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.User{}, models.SecurityQuestions{}, models.AnnuallyPlan{}, models.AnnuallyGoal{},
		models.MonthlyPlan{}, models.MonthlyGoal{},
		models.DailyPlan{}, models.DailyGoal{},
		models.Note{})

	db, err = gorm.Open(mysql.Open(source), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.User{}, models.SecurityQuestions{}, models.AnnuallyPlan{}, models.AnnuallyGoal{},
		models.MonthlyPlan{}, models.MonthlyGoal{},
		models.DailyPlan{}, models.DailyGoal{},
		models.Note{})
}
