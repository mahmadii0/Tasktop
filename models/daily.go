package models

import (
	"Tasktop/configure"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type DailyPlan struct {
	gorm.Model
	DPID     int       `gorm:"primaryKey;autoIncrement" json:"DPId"`
	Progress int       `gorm:"not null" json:"progress"`
	Status   bool      `gorm:"not null" json:"status"`
	Date     time.Time `gorm:"type:date;not null" json:"date"`
	UserName string    `gorm:"not null" json:"username"` //Foregin-key
	User     User      `gorm:"foreignKey:UserName"`
}
type DailyGoal struct {
	gorm.Model
	DGID  int    `gorm:"primaryKey;autoIncrement" json:"DGId"`
	Title string `gorm:"not null;size:100" json:"title"`
	//TimeTD  is stand of Time To Do for this task
	TimeTD      time.Time   `json:"timeToDo"`
	Priority    string      `json:"priority;size:6"`
	Status      bool        `gorm:"not null" json:"status"`
	DPID        int         `gorm:"not null" json:"DPId"` //Foregin-key
	MGID        int         `gorm:"not null" json:"MGId"` //Foregin-key
	DailyPlan   DailyPlan   `gorm:"foreignKey:DPID"`
	MonthlyGoal MonthlyGoal `gorm:"foreignKey:MGID"`
}

//Daily Plan Functions
//
//func GetDailyPs(username string) ([]*DailyPlan, error) {
//	dps := make([]*DailyPlan, 0)
//	query := `SELECT * FROM dailyplans WHERE username=?`
//	rows, err := db.Query(query, username)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//	for rows.Next() {
//		dp := new(DailyPlan)
//		if err := rows.Scan(&dp.DPID, &dp.Progress, &dp.Status, &dp.Date, &dp.UserName); err != nil {
//			return nil, err
//		}
//		dps = append(dps, dp)
//	}
//	if err := rows.Err(); err != nil {
//		return nil, err
//	}
//	return dps, err
//}
//
//func GetDailyPId(username string, date string) int {
//	var id int = 0
//	var s int = 0
//	query := `SELECT dailyPId,status FROM dailyplans WHERE username=? and date=?`
//	row := db.QueryRow(query, username, date)
//	err := row.Scan(&id, &s)
//	if err != nil {
//		fmt.Printf("Error while fetch data:", err)
//	}
//	if s == 0 && id != 0 {
//		query := `UPDATE dailyplans SET status = 1 WHERE username=? and date=?`
//		_, err := db.Exec(query, username, date)
//		if err != nil {
//			fmt.Printf("Error while activate dailyplan: ", err)
//		}
//	}
//	return id
//}
//
//func GetDailyPById(dailyPId int) (*DailyPlan, error) {
//	var dailyP *DailyPlan
//	query := `SELECT * FROM dailyplans WHERE dailyPId=?`
//	row := db.QueryRow(query, dailyPId)
//	err := row.Scan(&dailyP.DPID, &dailyP.Status, &dailyP.UserName)
//	return dailyP, err
//}
//
//func AddDailyP(username string, date string) bool {
//	status := true
//	query := `INSERT INTO dailyplans(progress,status,date,username) VALUES (0,1,?,?)`
//	_, err := db.Exec(query, date, username)
//	if err != nil {
//		status = false
//	}
//	return status
//
//}
//
//func UpdateDailyP(dailyP *DailyPlan) bool {
//	status := true
//	query := `UPDATE dailyplans SET status=?, userId=? WHERE dailyPId=?`
//	_, err := db.Exec(query, dailyP.Status, dailyP.UserName, dailyP.DPID)
//	if err != nil {
//		status = false
//	}
//	return status
//}
//
//func DeleteDailyPlan(dailyPId int) bool {
//	status := true
//	query := `DELETE FROM dailyplans WHERE dailyPId=?`
//	_, err := db.Exec(query, dailyPId)
//	if err != nil {
//		status = false
//	}
//	return status
//}
//
////Daily Goal Function
//
//func GetDailyGStatuses(id int) ([]int, error) {
//	statuses := make([]int, 0)
//	query := `SELECT status FROM dailygoals WHERE dailyPId=?`
//	rows, err := db.Query(query, id)
//	if err != nil {
//		return nil, err
//	}
//	defer db.Close()
//	for rows.Next() {
//		i := 0
//		if err := rows.Scan(&i); err != nil {
//			return nil, err
//		}
//		statuses = append(statuses, i)
//	}
//	return statuses, nil
//}
//func GetDailyGs(id int) ([]*DailyGoal, error) {
//	dgs := make([]*DailyGoal, 0)
//	query := `SELECT * FROM dailygoals WHERE dailyPId=?`
//	rows, err := db.Query(query, id)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//	for rows.Next() {
//		dg := new(DailyGoal)
//		if err := rows.Scan(&dg.DGID, &dg.Title, &dg.TimeTD, &dg.Priority, &dg.Status, &dg.DPID, &dg.MGID); err != nil {
//			return nil, err
//		}
//		dgs = append(dgs, dg)
//	}
//	if err := rows.Err(); err != nil {
//		return nil, err
//	}
//	return dgs, err
//
//}
//
//func GetDailyGById(dailyGId int) (*DailyGoal, error) {
//	dailyGoal := &DailyGoal{}
//	query := `SELECT * FROM dailygoals WHERE dailyGId=?`
//	row := db.QueryRow(query, dailyGId)
//	err := row.Scan(&dailyGoal.DGID, &dailyGoal.Title, &dailyGoal.TimeTD, &dailyGoal.Priority,
//		&dailyGoal.Status, &dailyGoal.DPID, &dailyGoal.MGID)
//	if err != nil {
//		return nil, err
//	}
//	dailyGoal.Priority = dailyGoal.Priority[1:]
//	return dailyGoal, nil
//}
//
//func AddDailyG(dailyGoal *DailyGoal) bool {
//	status := true
//	query := `INSERT INTO dailygoals(title,timeToDo,priority,status,dailyPId,monthlyGId) VALUES (?,?,?,?,?,?)`
//	_, err := db.Exec(query, dailyGoal.Title, dailyGoal.TimeTD, dailyGoal.Priority, 1, dailyGoal.DPID, dailyGoal.MGID)
//	if err != nil {
//		status = false
//	}
//	return status
//}
//
//func UpdateDailyG(dailyGoal *DailyGoal) bool {
//	status := true
//	query := `UPDATE dailygoals SET title=?, timeToDo=?, priority=?, status=?, dailyPId=?, monthlyGId=? WHERE dailyGId=?`
//	_, err := db.Exec(query, dailyGoal.Title, dailyGoal.TimeTD, dailyGoal.Priority, dailyGoal.Status, dailyGoal.DPID, dailyGoal.MGID, dailyGoal.DGID)
//	if err != nil {
//		status = false
//	}
//	return status
//}
//
//func DeleteDailyG(dailyGId int) bool {
//	status := true
//	query := `DELETE FROM dailygoals WHERE dailyGId=?`
//	_, err := db.Exec(query, dailyGId)
//	if err != nil {
//		status = false
//	}
//	return status
//}
