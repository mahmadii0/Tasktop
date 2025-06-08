package configure

import (
	"Tasktop"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	db *sql.DB
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func CreateTables() {
	source := Tasktop.DATABASE_SOURCE
	d, err := sql.Open("mysql", source)
	if err != nil {
		log.Fatal(err)
	}
	user := `
	CREATE TABLE IF NOT EXISTS users(
	    userId INTEGER PRIMARY KEY AUTO_INCREMENT,
	    name varchar(150) NOT NULL,
	    email varchar(250) UNIQUE NOT NULL,
	    phone varchar(13) UNIQUE,
	    password varchar(400) NOT NULL
	    );`
	annuallyPlan := `
	CREATE TABLE IF NOT EXISTS annuallyPlans(
	    annuallyPId INTEGER PRIMARY KEY AUTO_INCREMENT,
	    status TINYINT NOT NULL,
	    year INTEGER NOT NULL,
	    userId INTEGER NOT NULL,
	    CONSTRAINT user_annuallyP FOREIGN KEY (userId) REFERENCES users(userId)
	);`
	annuallyGoals := `
	CREATE TABLE IF NOT EXISTS annuallyGoals(
	    annuallyGId INTEGER PRIMARY KEY AUTO_INCREMENT,
	    title varchar(100) NOT NULL,
	    description varchar(1500),
	    status TINYINT NOT NULL,
	    annuallyPId INTEGER NOT NULL,
	    CONSTRAINT annuallyP_annuallyG FOREIGN KEY (annuallyPId) REFERENCES annuallyPlans(annuallyPId)
	);`
	monthlyPlan := `
	CREATE TABLE IF NOT EXISTS monthlyPlans(
	    monthlyPId INTEGER PRIMARY KEY AUTO_INCREMENT,
	    status TINYINT NOT NULL,
	    date varchar(40) NOT NULL,
	    userId INTEGER NOT NULL,
	    CONSTRAINT user_monthlyP FOREIGN KEY (userId) REFERENCES users(userId)
	);`
	monthlyGoals := `
	CREATE TABLE IF NOT EXISTS monthlyGoals(
	    monthlyGId INTEGER PRIMARY KEY AUTO_INCREMENT,
	    title varchar(100) NOT NULL,
	    description varchar(1500),
	    status TINYINT NOT NULL,
	    monthlyPId INTEGER NOT NULL,
	    annuallyGId INTEGER NOT NULL,
	    CONSTRAINT monthlyP_monthlyG FOREIGN KEY (monthlyPId) REFERENCES monthlyPlans(monthlyPId),
	    CONSTRAINT annuallyG_monthlyG FOREIGN KEY (annuallyGId) REFERENCES annuallyGoals(annuallyGId)
	);`
	dailyPlan := `
	CREATE TABLE IF NOT EXISTS dailyPlans(
	    dailyPId INTEGER PRIMARY KEY AUTO_INCREMENT,
	    status TINYINT NOT NULL,
	    userId INTEGER NOT NULL,
	    CONSTRAINT user_dailyP FOREIGN KEY (userId) REFERENCES users(userId)
	);`
	dailyGoals := `
	CREATE TABLE IF NOT EXISTS dailyGoals(
	    dailyGId INTEGER PRIMARY KEY AUTO_INCREMENT,
	    title varchar(100) NOT NULL,
	    timeToDo DATETIME,
	    status TINYINT NOT NULL,
	    dailyPId INTEGER NOT NULL,
	    monthlyGId INTEGER NOT NULL,
	    CONSTRAINT dailyP_dailyG FOREIGN KEY (dailyPId) REFERENCES dailyPlans(dailyPId),
	    CONSTRAINT monthlyG_dailyG FOREIGN KEY (monthlyGId) REFERENCES monthlyGoals(monthlyGId)
	);`

	if _, err = d.Exec(user); err != nil {
		log.Printf("Error on creating user table: %v", err)
	}
	if _, err = d.Exec(annuallyPlan); err != nil {
		log.Printf("Error on creating annuallyPlan table: %v", err)
	}
	if _, err = d.Exec(annuallyGoals); err != nil {
		log.Printf("Error on creating annuallyGoals table: %v", err)
	}
	if _, err = d.Exec(monthlyPlan); err != nil {
		log.Printf("Error on creating monthlyPlan table: %v", err)
	}
	if _, err = d.Exec(monthlyGoals); err != nil {
		log.Printf("Error on creating monthlyGoals table: %v", err)
	}
	if _, err = d.Exec(dailyPlan); err != nil {
		log.Printf("Error on creating dailyPlan table: %v", err)
	}
	if _, err = d.Exec(dailyGoals); err != nil {
		log.Printf("Error on creating dailyGoals table: %v", err)
	}

}

func Connect() {
	source := os.Getenv("DATA_SOURCE_NAME")
	d, err := sql.Open("mysql", source)
	if err != nil {
		log.Fatal(err)
	}
	db = d
}

func GetDB() *sql.DB {
	return db
}
