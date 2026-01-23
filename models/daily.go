package models

import (
	"Tasktop/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type DailyPlan struct {
	DPID     int       `gorm:"primaryKey;autoIncrement" json:"DPId"`
	Progress int       `gorm:"not null" json:"progress"`
	Status   bool      `gorm:"not null" json:"status"`
	Date     time.Time `gorm:"type:date;not null" json:"date"`
	UserID   int64     `gorm:"not null;index" json:"userId"` //Foregin-key
	User     User      `gorm:"foreignKey:UserID;references:ID"`
}
type DailyGoal struct {
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

func GetDailyPs(userId int64) ([]DailyPlan, error) {
	dp, err := gorm.G[DailyPlan](db).Where("user_id=?", userId).Find(ctx)
	if err != nil {
		return nil, err
	}
	return dp, nil
}

func GetDailyPId(userId int64, date string) int {
	id := 0
	ap, err := gorm.G[DailyPlan](db).Where("user_id=? and date=?", userId, date).Find(ctx)
	if err != nil {
		fmt.Printf("Error while fetch data:", err)
	}
	id = id + ap[0].DPID
	return id
}

//func GetDailyPById(dailyPId int) (*DailyPlan, error) {
//	var dailyP *DailyPlan
//	query := `SELECT * FROM dailyplans WHERE dailyPId=?`
//	row := db.QueryRow(query, dailyPId)
//	err := row.Scan(&dailyP.DPID, &dailyP.Status, &dailyP.UserName)
//	return dailyP, err
//}

func AddDailyP(userId int64, date string) bool {
	time, err := utils.ParseTime("date only", date)
	if err != nil {
		return false
	}
	gorm.G[DailyPlan](db).Create(ctx, &DailyPlan{
		Progress: 0,
		Status:   true,
		Date:     time,
		UserID:   userId,
	})
	return true
}

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

//Daily Goal Function

func GetDailyGStatuses(id int) ([]bool, error) {
	statuses := make([]bool, 0)
	dg, err := gorm.G[DailyGoal](db).Where("dp_id=?", id).Select("status").Find(ctx)
	if err != nil {
		return nil, err
	}
	for _, item := range dg {
		statuses = append(statuses, item.Status)
	}
	return statuses, nil
}
func GetDailyGs(id int) ([]DailyGoal, error) {
	dg, err := gorm.G[DailyGoal](db).Where("dp_id=?", id).Find(ctx)
	if err != nil {
		return nil, err
	}
	return dg, err

}

func GetDailyGById(dailyGId int) (DailyGoal, error) {
	dg, err := gorm.G[DailyGoal](db).Where("dg_id=?", dailyGId).Find(ctx)
	if err != nil {
		return dg[0], err
	}
	return dg[0], nil
}

func AddDailyG(dailyGoal *DailyGoal) bool {
	gorm.G[DailyGoal](db).Create(ctx, dailyGoal)
	return true
}

func UpdateDailyG(dailyGoal *DailyGoal) bool {
	gorm.G[DailyGoal](db).Where("dg_id=?", dailyGoal.DGID).Updates(ctx, *dailyGoal)
	return true
}

func DeleteDailyG(dailyGId int) bool {
	gorm.G[DailyGoal](db).Where("dg_id=?", dailyGId).Delete(ctx)
	return true
}
