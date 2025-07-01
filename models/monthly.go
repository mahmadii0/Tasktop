package models

import (
	"Tasktop/configure"
)

type MonthlyPlan struct {
	MPID     int    `json:"MPId"`
	Progress int    `json:"progress"`
	Status   bool   `json:"status"`
	Date     string `json:"date"`
	UserID   int    `json:"userId"` //Foregin-key
}
type MonthlyGoal struct {
	MGID     int    `json:"MGId"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Priority string `json:"priority"`
	Progress int    `json:"progress"`
	Status   bool   `json:"status"`
	MPID     int    `json:"MPId"` //Foregin-key
	AGID     int    `json:"AGId"` //Foregin-key
}

func init() {
	configure.Connect()
	db = configure.GetDB()
}

//Monthly Plan Functions

func GetMonthlyPByUserId(userId int) (*MonthlyPlan, error) {
	var monthlyP *MonthlyPlan
	query := `SELECT * FROM monthlyPlans WHERE userId=?`
	row := db.QueryRow(query, userId)
	err := row.Scan(&monthlyP.MPID, &monthlyP.Status, &monthlyP.Date, &monthlyP.UserID)
	return monthlyP, err
}

func GetMonthlyPById(monthlyPId int) (*MonthlyPlan, error) {
	var monthlyP *MonthlyPlan
	query := `SELECT * FROM monthlyPlans WHERE monthlyPId=?`
	row := db.QueryRow(query, monthlyPId)
	err := row.Scan(&monthlyP.MPID, &monthlyP.Status, &monthlyP.Date, &monthlyP.UserID)
	return monthlyP, err
}

func AddMonthlyP(date string, userId int) bool {
	status := true
	query := `INSERT INTO monthlyPlans(status,date,userId) VALUES (0,?,?)`
	_, err := db.Exec(query, date, userId)
	if err != nil {
		status = false
	}
	return status

}

func UpdateMonthlyP(monthlyP *MonthlyPlan) bool {
	status := true
	query := `UPDATE monthlyPlans SET status=?, date=?, userId=? WHERE monthlyPId=?`
	_, err := db.Exec(query, monthlyP.Status, monthlyP.Date, monthlyP.UserID, monthlyP.MPID)
	if err != nil {
		status = false
	}
	return status
}

func DeleteMonthlyPlan(monthlyPId int) bool {
	status := true
	query := `DELETE FROM monthlyPlans WHERE monthlyPId=?`
	_, err := db.Exec(query, monthlyPId)
	if err != nil {
		status = false
	}
	return status
}

//Monthly Goal Function

func GetMonthlyGByMonthlyPId(monthlyPId int) (*MonthlyGoal, error) {
	var monthlyGoal *MonthlyGoal
	query := `SELECT * FROM monthlyGoals WHERE monthlyGId=?`
	row := db.QueryRow(query, monthlyPId)
	err := row.Scan(&monthlyGoal.MGID, &monthlyGoal.Title, &monthlyGoal.Desc,
		&monthlyGoal.Status, &monthlyGoal.MPID, &monthlyGoal.AGID)
	return monthlyGoal, err
}

func GetMonthlyGById(monthlyGId int) (*MonthlyGoal, error) {
	var monthlyGoal *MonthlyGoal
	query := `SELECT * FROM annuallyGoals WHERE annuallyGId=?`
	row := db.QueryRow(query, monthlyGId)
	err := row.Scan(&monthlyGoal.MGID, &monthlyGoal.Title, &monthlyGoal.Desc,
		&monthlyGoal.Status, &monthlyGoal.MPID, &monthlyGoal.AGID)
	return monthlyGoal, err
}

func AddMonthlyG(monthlyGoal *MonthlyGoal) bool {
	status := true
	query := `INSERT INTO monthlyGoals(title,description,status,annuallyGId) VALUES (?,?,?,?)`
	_, err := db.Exec(query, monthlyGoal.Title, monthlyGoal.Desc, monthlyGoal.Status, monthlyGoal.AGID)
	if err != nil {
		status = false
	}
	return status
}

func UpdateMonthlyG(monthlyGoal *MonthlyGoal) bool {
	status := true
	query := `UPDATE monthlyGoals SET title=?, description=?, status=?, annuallyGId=? WHERE monthlyGId=?`
	_, err := db.Exec(query, monthlyGoal.Title, monthlyGoal.Desc, monthlyGoal.Status, monthlyGoal.AGID, monthlyGoal.MGID)
	if err != nil {
		status = false
	}
	return status
}

func DeleteMonthlyG(monthlyGId int) bool {
	status := true
	query := `DELETE FROM monthlyGoals WHERE monthlyGId=?`
	_, err := db.Exec(query, monthlyGId)
	if err != nil {
		status = false
	}
	return status
}
