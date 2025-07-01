package models

import (
	"Tasktop/configure"
	"time"
)

type DailyPlan struct {
	DPID     int  `json:"DPId"`
	Progress int  `json:"progress"`
	Status   bool `json:"status"`
	UserID   int  `json:"userId"` //Foregin-key
}
type DailyGoal struct {
	DGID  int    `json:"DGId"`
	Title string `json:"title"`
	//TimeTD  is stand of Time To Do for this task
	TimeTD   time.Time `json:"timeToDo"`
	Priority string    `json:"priority"`
	Status   bool      `json:"status"`
	DPID     int       `json:"DPId"` //Foregin-key
	MGID     int       `json:"MGId"` //Foregin-key
}

func init() {
	configure.Connect()
	db = configure.GetDB()
}

//Daily Plan Functions

func GetDailyPByUserId(userId int) (*DailyPlan, error) {
	var dailyP *DailyPlan
	query := `SELECT * FROM dailyplans WHERE userId=?`
	row := db.QueryRow(query, userId)
	err := row.Scan(&dailyP.DPID, &dailyP.Status, &dailyP.UserID)
	return dailyP, err
}

func GetDailyPById(dailyPId int) (*DailyPlan, error) {
	var dailyP *DailyPlan
	query := `SELECT * FROM dailyplans WHERE dailyPId=?`
	row := db.QueryRow(query, dailyPId)
	err := row.Scan(&dailyP.DPID, &dailyP.Status, &dailyP.UserID)
	return dailyP, err
}

func AddDailyP(date string, userId int) bool {
	status := true
	query := `INSERT INTO dailyplans(status,userId) VALUES (0,?)`
	_, err := db.Exec(query, userId)
	if err != nil {
		status = false
	}
	return status

}

func UpdateDailyP(dailyP *DailyPlan) bool {
	status := true
	query := `UPDATE dailyplans SET status=?, userId=? WHERE dailyPId=?`
	_, err := db.Exec(query, dailyP.Status, dailyP.UserID, dailyP.DPID)
	if err != nil {
		status = false
	}
	return status
}

func DeleteDailyPlan(dailyPId int) bool {
	status := true
	query := `DELETE FROM dailyplans WHERE dailyPId=?`
	_, err := db.Exec(query, dailyPId)
	if err != nil {
		status = false
	}
	return status
}

//Daily Goal Function

func GetDailyGByDailyPId(dailyPId int) (*DailyGoal, error) {
	var dailyGoal *DailyGoal
	query := `SELECT * FROM dailygoals WHERE dailyGId=?`
	row := db.QueryRow(query, dailyPId)
	err := row.Scan(&dailyGoal.DGID, &dailyGoal.Title, &dailyGoal.TimeTD,
		&dailyGoal.Status, &dailyGoal.DPID, &dailyGoal.MGID)
	return dailyGoal, err
}

func GetDailyGById(dailyGId int) (*DailyGoal, error) {
	var dailyGoal *DailyGoal
	query := `SELECT * FROM dailygoals WHERE dailyGId=?`
	row := db.QueryRow(query, dailyGId)
	err := row.Scan(&dailyGoal.DGID, &dailyGoal.Title, &dailyGoal.TimeTD,
		&dailyGoal.Status, &dailyGoal.DPID, &dailyGoal.MGID)
	return dailyGoal, err
}

func AddDailyG(dailyGoal *DailyGoal) bool {
	status := true
	query := `INSERT INTO dailygoals(title,timeToDo,priority,status,dailyPId,monthlyGId) VALUES (?,?,?,?,?,?)`
	_, err := db.Exec(query, dailyGoal.Title, dailyGoal.TimeTD, dailyGoal.Priority, dailyGoal.Status, dailyGoal.DPID, dailyGoal.MGID)
	if err != nil {
		status = false
	}
	return status
}

func UpdateDailyG(dailyGoal *DailyGoal) bool {
	status := true
	query := `UPDATE dailygoals SET title=?, timeToDo=?, status=?, dailyPId=?, monthlyGId=? WHERE dailyGId=?`
	_, err := db.Exec(query, dailyGoal.Title, dailyGoal.TimeTD, dailyGoal.Status, dailyGoal.DPID, dailyGoal.MGID, dailyGoal.DGID)
	if err != nil {
		status = false
	}
	return status
}

func DeleteDailyG(dailyGId int) bool {
	status := true
	query := `DELETE FROM dailygoals WHERE dailyGId=?`
	_, err := db.Exec(query, dailyGId)
	if err != nil {
		status = false
	}
	return status
}
