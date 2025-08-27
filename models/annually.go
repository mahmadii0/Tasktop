package models

import (
	"Tasktop/configure"
	"fmt"
)

type AnnuallyPlan struct {
	APID     int    `json:"APId"`
	Progress int    `json:"progress"`
	Status   bool   `json:"status"`
	Year     int    `json:"year"`
	UserName string `json:"username"` //Foregin-key
}
type AnnuallyGoal struct {
	AGID     int    `json:"AGId"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Priority string `json:"priority"`
	Progress int    `json:"progress"`
	Status   bool   `json:"status"`
	APID     int    `json:"APId"` //Foregin-key
}

func init() {
	configure.Connect()
	db = configure.GetDB()
}

//Annually Plan Functions

func GetAnnuallyPs(username string) ([]*AnnuallyPlan, error) {
	aps := make([]*AnnuallyPlan, 0)
	query := `SELECT * FROM annuallyplans WHERE username=?`
	rows, err := db.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		ap := new(AnnuallyPlan)
		if err := rows.Scan(&ap.APID, &ap.Progress, &ap.Status, &ap.Year, &ap.UserName); err != nil {
			return nil, err
		}
		aps = append(aps, ap)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return aps, err
}

func GetAnnuallyPId(username string, year int) int {
	var id int = 0
	var s int = 0
	query := `SELECT annuallyPId,status FROM annuallyplans WHERE username=? and year=?`
	row := db.QueryRow(query, username, year)
	err := row.Scan(&id, &s)
	if err != nil {
		fmt.Printf("Error while fetch data:", err)
	}
	if s == 0 && id != 0 {
		query := `UPDATE annuallyplans SET status = 1 WHERE username=? and year=?`
		_, err := db.Exec(query, username, year)
		if err != nil {
			fmt.Printf("Error while activate annuallyplan: ", err)
		}
	}
	return id
}

func AddAnnuallyP(username string, year int) bool {
	status := true
	query := `INSERT INTO annuallyplans(progress,status,year,username) VALUES (0,1,?,?)`
	_, err := db.Exec(query, year, username)
	if err != nil {
		status = false
	}
	return status

}

//Annually Goal Function

func GetAnnuallyGIdByMonthlyGId(id int) int {
	var annuallyGId int
	query := `SELECT annuallyGId FROM monthlygoals WHERE monthlyGId=?`
	row := db.QueryRow(query, id)
	err := row.Scan(&annuallyGId)
	if err != nil {
		return -1
	}
	return annuallyGId
}

func GetAnnuallyGs(id int) ([]*AnnuallyGoal, error) {
	ags := make([]*AnnuallyGoal, 0)
	query := `SELECT * FROM annuallygoals WHERE annuallyPId=?`
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		ag := new(AnnuallyGoal)
		if err := rows.Scan(&ag.AGID, &ag.Desc, &ag.Title, &ag.Priority, &ag.Progress, &ag.Status, &ag.APID); err != nil {
			return nil, err
		}
		ags = append(ags, ag)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ags, err

}

func GetAnnuallyGById(annuallyGId int) (*AnnuallyGoal, error) {
	annuallyGoal := &AnnuallyGoal{}
	query := `SELECT * FROM annuallyGoals WHERE annuallyGId=?`
	row := db.QueryRow(query, annuallyGId)
	err := row.Scan(&annuallyGoal.AGID, &annuallyGoal.Title, &annuallyGoal.Desc, &annuallyGoal.Priority, &annuallyGoal.Progress,
		&annuallyGoal.Status, &annuallyGoal.APID)
	return annuallyGoal, err
}

func AddAnnuallyG(annuallyGoal *AnnuallyGoal) bool {
	status := true
	query := `INSERT INTO annuallyGoals(title,description,priority,progress,status,annuallyPId) VALUES (?,?,?,0,1,?)`
	_, err := db.Exec(query, annuallyGoal.Title, annuallyGoal.Desc, annuallyGoal.Priority, annuallyGoal.APID)
	if err != nil {
		status = false
	}
	return status
}

func UpdateAnnuallyG(annuallyGoal *AnnuallyGoal) bool {
	status := true
	query := `UPDATE annuallyGoals SET title=?, description=?, priority=?, progress=?, status=?, annuallyPId=? WHERE annuallyGId=?`
	_, err := db.Exec(query, annuallyGoal.Title, annuallyGoal.Desc, annuallyGoal.Priority, annuallyGoal.Progress, annuallyGoal.Status, annuallyGoal.APID, annuallyGoal.AGID)
	if err != nil {
		status = false
	}
	return status
}

func DeleteAnnuallyG(annuallyGId int) bool {
	status := true
	query := `DELETE FROM annuallyGoals WHERE annuallyGId=?`
	_, err := db.Exec(query, annuallyGId)
	if err != nil {
		status = false
	}
	return status
}

//Annually Report

func GetAProgresses(annuallyPs []*AnnuallyPlan) (map[int]int, error) {
	var progresses = make(map[int]int)
	for _, annuallyP := range annuallyPs {
		query := `SELECT progress FROM annuallygoals WHERE annuallyPId=?`
		rows, err := db.Query(query, annuallyP.APID)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var progress int
		counter := 0
		for rows.Next() {
			var p int
			if err := rows.Scan(&p); err != nil {
				return nil, err
			}
			progress = progress + p
			counter++
		}
		progress = progress / counter
		progresses[annuallyP.Year] = progress
	}
	return progresses, nil
}
