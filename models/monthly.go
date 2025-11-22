package models

import (
	"Tasktop/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type MonthlyPlan struct {
	gorm.Model
	MPID     int       `gorm:"primaryKey;autoIncrement" json:"MPId"`
	Progress int       `gorm:"not null" json:"progress"`
	Status   bool      `gorm:"not null" json:"status"`
	Date     time.Time `gorm:"type:date;not null" json:"date"`
	UserName string    `gorm:"not null;size:100" json:"username"` //Foregin-key
	User     User      `gorm:"foreignkey:UserName"`
}
type MonthlyGoal struct {
	gorm.Model
	MGID         int          `gorm:"primaryKey;autoIncrement" json:"MGId"`
	Title        string       `gorm:"not null;size:100" json:"title"`
	Desc         string       `gorm:"size:1500" json:"desc"`
	Priority     string       `gorm:"size:6" json:"priority"`
	Progress     int          `gorm:"not null" json:"progress"`
	Status       bool         `gorm:"not null" json:"status"`
	MPID         int          `gorm:"not null" json:"MPId"` //Foregin-key
	AGID         int          `gorm:"not null" json:"AGId"` //Foregin-key
	MonthlyPlan  MonthlyPlan  `gorm:"foreignkey:MPID"`
	AnnuallyGoal AnnuallyGoal `gorm:"foreignkey:AGID"`
}

//Monthly Plan Function

func GetMonthlyPs(username string) ([]MonthlyPlan, error) {
	mp, err := gorm.G[MonthlyPlan](db).Where("username=?", username).Find(ctx)
	if err != nil {
		return nil, err
	}
	return mp, nil
}

func GetMonthlyPId(username string, date string) int {
	id := 0
	ap, err := gorm.G[MonthlyPlan](db).Where("username=? and date=?", username, date).Find(ctx)
	if err != nil {
		fmt.Printf("Error while fetch data:", err)
	}
	id = id + ap[0].MPID
	return id
}

//func GetMonthlyPById(monthlyPId int) (*MonthlyPlan, error) {
//	var monthlyP *MonthlyPlan
//	query := `SELECT * FROM monthlyPlans WHERE monthlyPId=?`
//	row := db.QueryRow(query, monthlyPId)
//	err := row.Scan(&monthlyP.MPID, &monthlyP.Status, &monthlyP.Date, &monthlyP.UserName)
//	return monthlyP, err
//}

func AddMonthlyP(username string, date string) bool {
	time, err := utils.ParseTime("date only", date)
	if err != nil {
		return false
	}
	gorm.G[MonthlyPlan](db).Create(ctx, &MonthlyPlan{
		Progress: 0,
		Status:   true,
		Date:     time,
		UserName: username,
	})
	return true

}

//func UpdateMonthlyP(monthlyP *MonthlyPlan) bool {
//	status := true
//	query := `UPDATE monthlyPlans SET status=?, date=?, userId=? WHERE monthlyPId=?`
//	_, err := db.Exec(query, monthlyP.Status, monthlyP.Date, monthlyP.UserName, monthlyP.MPID)
//	if err != nil {
//		status = false
//	}
//	return status
//}

//func DeleteMonthlyPlan(monthlyPId int) bool {
//	status := true
//	query := `DELETE FROM monthlyPlans WHERE monthlyPId=?`
//	_, err := db.Exec(query, monthlyPId)
//	if err != nil {
//		status = false
//	}
//	return status
//}

//Monthly Goal Function

func GetMProgresses(monthlyPs []*MonthlyPlan) (map[time.Time]int, error) {
	var progresses = make(map[time.Time]int)
	for _, monthlyP := range monthlyPs {
		ag, err := gorm.G[MonthlyGoal](db).Where("monthlyPId=?", monthlyP.MPID).Select("progress").Find(ctx)
		if err != nil {
			return nil, err
		}
		var progress int
		counter := 0
		for _, v := range ag {
			p := v.Progress
			progress = progress + p
			counter++
		}
		progress = progress / counter
		progresses[monthlyP.Date] = progress
	}
	return progresses, nil

}

func GetMonthlyGIdByDailyGId(id int) int {
	dg, err := gorm.G[DailyGoal](db).Where("dailyGId=?", id).Select("monthlyGId").Find(ctx)
	if err != nil {
		return -1
	}
	return dg[0].MGID
}

func GetMonthlyGs(id int) ([]MonthlyGoal, error) {
	mg, err := gorm.G[MonthlyGoal](db).Where("monthlyPId=?", id).Find(ctx)
	if err != nil {
		return nil, err
	}
	return mg, err

}

func GetMonthlyGById(monthlyGId int) (MonthlyGoal, error) {
	mg, err := gorm.G[MonthlyGoal](db).Where("monthlyGId=?", monthlyGId).Find(ctx)
	if err != nil {
		return mg[0], err
	}
	return mg[0], nil
}

func AddMonthlyG(monthlyGoal *MonthlyGoal) bool {
	monthlyGoal.Progress = 0
	monthlyGoal.Status = true
	gorm.G[MonthlyGoal](db).Create(ctx, monthlyGoal)
	return true
}

func UpdateMonthlyG(monthlyGoal *MonthlyGoal) bool {
	gorm.G[MonthlyGoal](db).Where("monthlyGId=?", monthlyGoal.MGID).Updates(ctx, *monthlyGoal)
	return true
}

func DeleteMonthlyG(monthlyGId int) bool {
	gorm.G[MonthlyGoal](db).Where("monthlyGId=?", monthlyGId).Delete(ctx)
	return true
}
