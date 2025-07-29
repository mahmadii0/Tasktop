package models

import (
	"Tasktop/configure"
	"fmt"
)

type MonthlyPlan struct {
	MPID     int    `json:"MPId"`
	Progress int    `json:"progress"`
	Status   bool   `json:"status"`
	Date     string `json:"date"`
	UserName string `json:"userId"` //Foregin-key
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

func GetMonthlyPId(username string, date string) int {
	var id int = 0
	var s int = 0
	query := `SELECT monthlyPId,status FROM monthlyPlans WHERE username=? and date=?`
	row := db.QueryRow(query, username, date)
	err := row.Scan(&id, &s)
	if err != nil {
		fmt.Printf("Error while fetch data:", err)
	}
	if s == 0 && id != 0 {
		query := `UPDATE monthlyPlans SET status = 1 WHERE username=? and date=?`
		_, err := db.Exec(query, username, date)
		if err != nil {
			fmt.Printf("Error while activate monthlyplan: ", err)
		}
	}
	return id
}

func GetMonthlyPById(monthlyPId int) (*MonthlyPlan, error) {
	var monthlyP *MonthlyPlan
	query := `SELECT * FROM monthlyPlans WHERE monthlyPId=?`
	row := db.QueryRow(query, monthlyPId)
	err := row.Scan(&monthlyP.MPID, &monthlyP.Status, &monthlyP.Date, &monthlyP.UserName)
	return monthlyP, err
}

func AddMonthlyP(username string, date string) bool {
	status := true
	query := `INSERT INTO monthlyPlans(progress,status,date,username) VALUES (0,1,?,?)`
	_, err := db.Exec(query, date, username)
	if err != nil {
		status = false
	}
	return status

}

func UpdateMonthlyP(monthlyP *MonthlyPlan) bool {
	status := true
	query := `UPDATE monthlyPlans SET status=?, date=?, userId=? WHERE monthlyPId=?`
	_, err := db.Exec(query, monthlyP.Status, monthlyP.Date, monthlyP.UserName, monthlyP.MPID)
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

// func GetMonthlyGByMonthlyPId(monthlyPId int) (*MonthlyGoal, error) {
// 	var monthlyGoal *MonthlyGoal
// 	query := `SELECT * FROM monthlygoals WHERE monthlyGId=?`
// 	row := db.QueryRow(query, monthlyPId)
// 	err := row.Scan(&monthlyGoal.MGID, &monthlyGoal.Title, &monthlyGoal.Desc,
// 		&monthlyGoal.Status, &monthlyGoal.MPID, &monthlyGoal.AGID)
// 	return monthlyGoal, err
// }

func GetMonthlyGs(id int) ([]*MonthlyGoal, error) {
	mgs := make([]*MonthlyGoal, 0)
	query := `SELECT * FROM monthlygoals WHERE monthlyPId=?`
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mg := new(MonthlyGoal)
		if err := rows.Scan(&mg.MGID, &mg.Title, &mg.Desc, &mg.Priority, &mg.Progress, &mg.Status, &mg.MPID, &mg.AGID); err != nil {
			return nil, err
		}
		mgs = append(mgs, mg)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return mgs, err

}

func GetMonthlyGById(monthlyGId int) (*MonthlyGoal, error) {
	monthlyGoal := &MonthlyGoal{}
	query := `SELECT * FROM monthlyGoals WHERE monthlyGId=?`
	row := db.QueryRow(query, monthlyGId)
	err := row.Scan(&monthlyGoal.MGID, &monthlyGoal.Title, &monthlyGoal.Desc, &monthlyGoal.Priority, &monthlyGoal.Progress,
		&monthlyGoal.Status, &monthlyGoal.MPID, &monthlyGoal.AGID)
	return monthlyGoal, err
}

func AddMonthlyG(monthlyGoal *MonthlyGoal) bool {
	status := true
	query := `INSERT INTO monthlygoals(title,description,priority,progress,status,monthlyPId,annuallyGId) VALUES (?,?,?,0,1,?,?)`
	_, err := db.Exec(query, monthlyGoal.Title, monthlyGoal.Desc, monthlyGoal.Priority, monthlyGoal.MPID, monthlyGoal.AGID)
	if err != nil {
		status = false
	}
	return status
}

func UpdateMonthlyG(monthlyGoal *MonthlyGoal) bool {
	status := true
	query := `UPDATE monthlygoals SET title=?, description=?, priority=?, progress=?, status=?, monthlyPId=?, annuallyGId=? WHERE monthlyGId=?`
	_, err := db.Exec(query, monthlyGoal.Title, monthlyGoal.Desc, monthlyGoal.Priority, monthlyGoal.Progress, monthlyGoal.Status, monthlyGoal.MPID, monthlyGoal.AGID, monthlyGoal.MGID)
	if err != nil {
		status = false
	}
	return status
}

func DeleteMonthlyG(monthlyGId int) bool {
	status := true
	query := `DELETE FROM monthlygoals WHERE monthlyGId=?`
	_, err := db.Exec(query, monthlyGId)
	if err != nil {
		status = false
	}
	return status
}
