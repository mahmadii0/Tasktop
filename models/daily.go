package models

import (
	"Tasktop/configure"
	"time"
)

type DailyPlan struct {
	DPID   int
	Status bool
	UserID int //Foregin-key
}
type DailyGoal struct {
	DGID  int
	Title string
	//TimeTD  is stand of Time To Do for this task
	TimeTD time.Time
	Status bool
	DPID   int //Foregin-key
	MGID   int //Foregin-key
}

func init() {
	configure.Connect()
	db = configure.GetDB()
}

//Daily Plan Functions

func GetDailyPByUserId(userId int) (*DailyPlan, error) {
	var dailyP *DailyPlan
	query := `SELECT * FROM dailyPlans WHERE userId=?`
	row := db.QueryRow(query, userId)
	err := row.Scan(&dailyP.DPID, &dailyP.Status, &dailyP.UserID)
	return dailyP, err
}

func GetDailyPById(dailyPId int) (*DailyPlan, error) {
	var dailyP *DailyPlan
	query := `SELECT * FROM dailyPlans WHERE dailyPId=?`
	row := db.QueryRow(query, dailyPId)
	err := row.Scan(&dailyP.DPID, &dailyP.Status, &dailyP.UserID)
	return dailyP, err
}

func AddDailyP(date string, userId int) bool {
	status := true
	query := `INSERT INTO dailyPlans(status,userId) VALUES (0,?)`
	_, err := db.Exec(query, userId)
	if err != nil {
		status = false
	}
	return status

}

func UpdateDailyP(dailyP *DailyPlan) bool {
	status := true
	query := `UPDATE dailyPlans SET status=?, userId=? WHERE dailyPId=?`
	_, err := db.Exec(query, dailyP.Status, dailyP.UserID, dailyP.DPID)
	if err != nil {
		status = false
	}
	return status
}

func DeleteDailyPlan(dailyPId int) bool {
	status := true
	query := `DELETE FROM dailyPlans WHERE dailyPId=?`
	_, err := db.Exec(query, dailyPId)
	if err != nil {
		status = false
	}
	return status
}

//Daily Goal Function

func GetDailyGByDailyPId(dailyPId int) (*DailyGoal, error) {
	var dailyGoal *DailyGoal
	query := `SELECT * FROM dailyGoals WHERE dailyGId=?`
	row := db.QueryRow(query, dailyPId)
	err := row.Scan(&dailyGoal.DGID, &dailyGoal.Title, &dailyGoal.TimeTD,
		&dailyGoal.Status, &dailyGoal.DPID, &dailyGoal.MGID)
	return dailyGoal, err
}

func GetDailyGById(dailyGId int) (*DailyGoal, error) {
	var dailyGoal *DailyGoal
	query := `SELECT * FROM dailyGoals WHERE dailyGId=?`
	row := db.QueryRow(query, dailyGId)
	err := row.Scan(&dailyGoal.DGID, &dailyGoal.Title, &dailyGoal.TimeTD,
		&dailyGoal.Status, &dailyGoal.DPID, &dailyGoal.MGID)
	return dailyGoal, err
}

func AddDailyG(dailyGoal *DailyGoal) bool {
	status := true
	query := `INSERT INTO dailyGoals(title,timeToDo,status,dailyPId,monthlyGId) VALUES (?,?,?,?,?)`
	_, err := db.Exec(query, dailyGoal.Title, dailyGoal.TimeTD, dailyGoal.Status, dailyGoal.DPID, dailyGoal.MGID)
	if err != nil {
		status = false
	}
	return status
}

func UpdateDailyG(dailyGoal *DailyGoal) bool {
	status := true
	query := `UPDATE dailyGoals SET title=?, timeToDo=?, status=?, dailyPId=?, monthlyGId=? WHERE dailyGId=?`
	_, err := db.Exec(query, dailyGoal.Title, dailyGoal.TimeTD, dailyGoal.Status, dailyGoal.DPID, dailyGoal.MGID, dailyGoal.DGID)
	if err != nil {
		status = false
	}
	return status
}

func DeleteDailyG(dailyGId int) bool {
	status := true
	query := `DELETE FROM dailyGoals WHERE dailyGId=?`
	_, err := db.Exec(query, dailyGId)
	if err != nil {
		status = false
	}
	return status
}
