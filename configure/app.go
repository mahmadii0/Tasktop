package configure

import (
	"Tasktop/constants"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func CreateTables() {
	source := constants.DATABASE_SOURCE
	d, err := sql.Open("mysql", source)
	if err != nil {
		log.Fatal(err)
	}
	user := `
	CREATE TABLE IF NOT EXISTS users(
	    username varchar(100) PRIMARY KEY,
	    fullname varchar(150) NOT NULL,
	    email varchar(250) UNIQUE NOT NULL,
	    phone varchar(13) UNIQUE,
	    password varchar(400) NOT NULL,
		session_token varchar(64),
		csrf_token varchar(64)
	    );`

	securityQuestions :=
		`CREATE TABLE IF NOT EXISTS securityquestions(
	    username varchar(100) NOT NULL,
	    question1 varchar(2) NOT NULL,
	    answer1 varchar(100) NOT NULL,
	    question2 varchar(2) NOT NULL,
	    answer2 varchar(100) NOT NULL,
		CONSTRAINT user_securityQ FOREIGN KEY (username) REFERENCES users(username)
	    );`

	annuallyPlan := `
	CREATE TABLE IF NOT EXISTS annuallyPlans(
	    annuallyPId INTEGER PRIMARY KEY AUTO_INCREMENT,
	    progress INTEGER NOT NULL,
	    status TINYINT NOT NULL,
	    year INTEGER NOT NULL,
	    username varchar(100) NOT NULL,
	    CONSTRAINT user_annuallyP FOREIGN KEY (username) REFERENCES users(username)
	);`
	annuallyGoals := `
	CREATE TABLE IF NOT EXISTS annuallyGoals(
	    annuallyGId INTEGER PRIMARY KEY AUTO_INCREMENT,
	    title varchar(100) NOT NULL,
	    description varchar(1500),
	    priority varchar(6),
	    progress INTEGER NOT NULL,
	    status TINYINT NOT NULL,
	    annuallyPId INTEGER NOT NULL,
	    CONSTRAINT annuallyP_annuallyG FOREIGN KEY (annuallyPId) REFERENCES annuallyPlans(annuallyPId)
	);`
	monthlyPlan := `
	CREATE TABLE IF NOT EXISTS monthlyPlans(
	    monthlyPId INTEGER PRIMARY KEY AUTO_INCREMENT,
	    progress INTEGER NOT NULL,
	    status TINYINT NOT NULL,
	    date DATE NOT NULL,
	    username varchar(100) NOT NULL,
	    CONSTRAINT user_monthlyP FOREIGN KEY (username) REFERENCES users(username)
	);`
	monthlyGoals := `
	CREATE TABLE IF NOT EXISTS monthlyGoals(
	    monthlyGId INTEGER PRIMARY KEY AUTO_INCREMENT,
	    title varchar(100) NOT NULL,
	    description varchar(1500),
	    priority varchar(6),
	    progress INTEGER NOT NULL,
	    status TINYINT NOT NULL,
	    monthlyPId INTEGER NOT NULL,
	    annuallyGId INTEGER NOT NULL,
	    CONSTRAINT monthlyP_monthlyG FOREIGN KEY (monthlyPId) REFERENCES monthlyPlans(monthlyPId),
	    CONSTRAINT annuallyG_monthlyG FOREIGN KEY (annuallyGId) REFERENCES annuallyGoals(annuallyGId)
	);`
	dailyPlan := `
	CREATE TABLE IF NOT EXISTS dailyPlans(
	    dailyPId INTEGER PRIMARY KEY AUTO_INCREMENT,
	    progress INTEGER NOT NULL,
	    status TINYINT NOT NULL,
		date DATE NOT NULL,
	    username varchar(100) NOT NULL,
	    CONSTRAINT user_dailyP FOREIGN KEY (username) REFERENCES users(username)
	);`
	dailyGoals := `
	CREATE TABLE IF NOT EXISTS dailyGoals(
	    dailyGId INTEGER PRIMARY KEY AUTO_INCREMENT,
	    title varchar(100) NOT NULL,
	    timeToDo DATETIME,
		priority varchar(6),
	    status TINYINT NOT NULL,
	    dailyPId INTEGER NOT NULL,
	    monthlyGId INTEGER NOT NULL,
	    CONSTRAINT dailyP_dailyG FOREIGN KEY (dailyPId) REFERENCES dailyPlans(dailyPId),
	    CONSTRAINT monthlyG_dailyG FOREIGN KEY (monthlyGId) REFERENCES monthlyGoals(monthlyGId)
	);`
	notes := `
	CREATE TABLE IF NOT EXISTS notes(
		noteId INTEGER PRIMARY KEY AUTO_INCREMENT,
		title varchar(200),
		note_text varchar(2000) NOT NULL,
		username varchar(100) NOT NULL,
		CONSTRAINT user_notes FOREIGN KEY (username) REFERENCES users(username)
	);`

	if _, err = d.Exec(user); err != nil {
		log.Printf("Error on creating user table: %v", err)
	}
	if _, err = d.Exec(securityQuestions); err != nil {
		log.Printf("Error on creating securityQuestions table: %v", err)
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
	if _, err = d.Exec(notes); err != nil {
		log.Printf("Error on creating notes table: %v", err)
	}

}

func Connect() {
	source := constants.DATABASE_SOURCE
	d, err := sql.Open("mysql", source)
	if err != nil {
		log.Fatal(err)
	}
	db = d
}

func GetDB() *sql.DB {
	return db
}
