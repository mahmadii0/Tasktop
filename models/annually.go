package models

import (
	"Tasktop/configure"
)

type AnnuallyPlan struct {
	APID     int
	Progress int
	Status   bool
	Year     int
	UserID   int //Foregin-key
}
type AnnuallyGoal struct {
	AGID     int
	Title    string
	Desc     string
	Priority string
	Progress int
	Status   bool
	APID     int //Foregin-key
}

func init() {
	configure.Connect()
	db = configure.GetDB()
}

//Annually Plan Functions

func GetAnnuallyPByUserId(userId int) (*AnnuallyPlan, error) {
	annuallyP := &AnnuallyPlan{}
	query := `SELECT * FROM annuallyPlans WHERE userId=?`
	row := db.QueryRow(query, userId)
	err := row.Scan(&annuallyP.APID, &annuallyP.Status, &annuallyP.Year, &annuallyP.UserID)
	return annuallyP, err
}
func GetAnnuallyPById(annuallyPId int) (*AnnuallyPlan, error) {
	var annuallyP *AnnuallyPlan
	query := `SELECT * FROM annuallyPlans WHERE annuallyPId=?`
	row := db.QueryRow(query, annuallyPId)
	err := row.Scan(&annuallyP.APID, &annuallyP.Status, &annuallyP.Year, &annuallyP.UserID)
	return annuallyP, err
}

func AddAnnuallyP(year int, userId int) bool {
	status := true
	query := `INSERT INTO annuallyPlans(status,year,userId) VALUES (0,?,?)`
	_, err := db.Exec(query, year, userId)
	if err != nil {
		status = false
	}
	return status

}

func UpdateAnnuallyP(annuallyP *AnnuallyPlan) bool {
	status := true
	query := `UPDATE annuallyPlans SET status=?, year=?, userId=? WHERE annuallyPId=?`
	_, err := db.Exec(query, annuallyP.Status, annuallyP.Year, annuallyP.UserID, annuallyP.APID)
	if err != nil {
		status = false
	}
	return status
}

func DeleteAnnuallyPlan(annuallyPId int) bool {
	status := true
	query := `DELETE FROM annuallyPlans WHERE annuallyPId=?`
	_, err := db.Exec(query, annuallyPId)
	if err != nil {
		status = false
	}
	return status
}

//Annually Goal Function

func GetAnnuallyGByAnnuallyPId(annuallyPId int) (*AnnuallyGoal, error) {
	var annuallyGoal *AnnuallyGoal
	query := `SELECT * FROM annuallyGoals WHERE annuallyPId=?`
	row := db.QueryRow(query, annuallyPId)
	err := row.Scan(&annuallyGoal.AGID, &annuallyGoal.Title, &annuallyGoal.Desc,
		&annuallyGoal.Status, &annuallyGoal.APID)
	return annuallyGoal, err
}

func GetAnnuallyGById(annuallyGId int) (*AnnuallyGoal, error) {
	var annuallyGoal *AnnuallyGoal
	query := `SELECT * FROM annuallyGoals WHERE annuallyGId=?`
	row := db.QueryRow(query, annuallyGId)
	err := row.Scan(&annuallyGoal.AGID, &annuallyGoal.Title, &annuallyGoal.Desc,
		&annuallyGoal.Status, &annuallyGoal.APID)
	return annuallyGoal, err
}

func AddAnnuallyG(annuallyGoal *AnnuallyGoal) bool {
	status := true
	query := `INSERT INTO annuallyGoals(title,description,status,annuallyPId) VALUES (?,?,?,?)`
	_, err := db.Exec(query, annuallyGoal.Title, annuallyGoal.Desc, annuallyGoal.Status, annuallyGoal.APID)
	if err != nil {
		status = false
	}
	return status
}

func UpdateAnnuallyG(annuallyGoal *AnnuallyGoal) bool {
	status := true
	query := `UPDATE annuallyGoals SET title=?, description=?, status=?, annuallyPId=? WHERE annuallyGId=?`
	_, err := db.Exec(query, annuallyGoal.Title, annuallyGoal.Desc, annuallyGoal.Status, annuallyGoal.APID, annuallyGoal.AGID)
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
