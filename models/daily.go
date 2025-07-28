package models

import (
	"Tasktop/configure"
	"fmt"
)

type DailyPlan struct {
	DPID     int    `json:"DPId"`
	Progress int    `json:"progress"`
	Status   bool   `json:"status"`
	UserName string `json:"userId"` //Foregin-key
}
type DailyGoal struct {
	DGID  int    `json:"DGId"`
	Title string `json:"title"`
	//TimeTD  is stand of Time To Do for this task
	TimeTD   string `json:"timeToDo"`
	Priority string `json:"priority"`
	Status   bool   `json:"status"`
	DPID     int    `json:"DPId"` //Foregin-key
	MGID     int    `json:"MGId"` //Foregin-key
}

func init() {
	configure.Connect()
	db = configure.GetDB()
}

//Daily Plan Functions

func GetDailyPId(username string, date string) int {
	var id int = 0
	var s int = 0
	query := `SELECT dailyPId,status FROM dailyplans WHERE username=? and date=?`
	row := db.QueryRow(query, username, date)
	err := row.Scan(&id, &s)
	if err != nil {
		fmt.Printf("Error while fetch data:", err)
	}
	if s == 0 && id != 0 {
		query := `UPDATE dailyplans SET status = 1 WHERE username=? and date=?`
		_, err := db.Exec(query, username, date)
		if err != nil {
			fmt.Printf("Error while activate dailyplan: ", err)
		}
	}
	return id
}

func GetDailyPById(dailyPId int) (*DailyPlan, error) {
	var dailyP *DailyPlan
	query := `SELECT * FROM dailyplans WHERE dailyPId=?`
	row := db.QueryRow(query, dailyPId)
	err := row.Scan(&dailyP.DPID, &dailyP.Status, &dailyP.UserName)
	return dailyP, err
}

func AddDailyP(username string, date string) bool {
	status := true
	query := `INSERT INTO dailyplans(progress,status,date,username) VALUES (0,1,?,?)`
	_, err := db.Exec(query, date, username)
	if err != nil {
		status = false
	}
	return status

}

func UpdateDailyP(dailyP *DailyPlan) bool {
	status := true
	query := `UPDATE dailyplans SET status=?, userId=? WHERE dailyPId=?`
	_, err := db.Exec(query, dailyP.Status, dailyP.UserName, dailyP.DPID)
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
	_, err := db.Exec(query, dailyGoal.Title, dailyGoal.TimeTD, dailyGoal.Priority, 1, dailyGoal.DPID, dailyGoal.MGID)
	if err != nil {
		status = false
	}
	return status
}

func UpdateDailyG(dailyGoal *DailyGoal) bool {
	status := true
	query := `UPDATE dailygoals SET title=?, timeToDo=?, priority=?, status=?, dailyPId=?, monthlyGId=? WHERE dailyGId=?`
	_, err := db.Exec(query, dailyGoal.Title, dailyGoal.TimeTD, dailyGoal.Priority, dailyGoal.Status, dailyGoal.DPID, dailyGoal.MGID, dailyGoal.DGID)
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
