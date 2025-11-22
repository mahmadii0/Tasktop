package models

import (
	"Tasktop/utils"
	"fmt"

	"gorm.io/gorm"
)

// status 1 is active
type AnnuallyPlan struct {
	gorm.Model
	APID     int    `gorm:"primaryKey;autoIncrement" json:"APId"`
	Progress int    `gorm:"not null" json:"progress"`
	Status   bool   `gorm:"not null" json:"status"`
	Year     int    `gorm:"not null" json:"year"`
	UserName string `gorm:"not null;size:100" json:"username"` //Foregin-key
	User     User   `gorm:"foreignKey:UserName"`
}
type AnnuallyGoal struct {
	gorm.Model
	AGID         int          `gorm:"primaryKey;autoIncrement" json:"AGId"`
	Title        string       `gorm:"not null;size:100" json:"title"`
	Desc         string       `gorm:"size:1500" json:"desc"`
	Priority     string       `gorm:"size:6" json:"priority"`
	Progress     int          `gorm:"not null" json:"progress"`
	Status       bool         `gorm:"not null" jgorm:"not null" son:"status"`
	APID         int          `gorm:"not null" json:"APId"` //Foregin-key
	AnnuallyPlan AnnuallyPlan `gorm:"foreignKey:APID"`
}

func init() {
	db, ctx = utils.GetDBctx()
}

//Annually Plan Functions

func GetAnnuallyPs(username string) ([]AnnuallyPlan, error) {
	ap, err := gorm.G[AnnuallyPlan](db).Where("username=?", username).Find(ctx)
	if err != nil {
		return nil, err
	}
	return ap, nil

}

func GetAnnuallyPId(username string, year int) int {
	id := 0
	ap, err := gorm.G[AnnuallyPlan](db).Where("username=? and year=?", username, year).Find(ctx)
	if err != nil {
		fmt.Printf("Error while fetch data:", err)
	}
	id = id + ap[0].APID
	return id
}

func AddAnnuallyP(username string, year int) bool {
	ap, err := gorm.G[AnnuallyPlan](db).Where("year = ? and username=?", year, username).Find(ctx)
	if err != nil {
		fmt.Printf("Error while fetch data:", err)
	}
	if len(ap) == 0 {
		gorm.G[AnnuallyPlan](db).Create(ctx, &AnnuallyPlan{
			Progress: 0,
			Status:   true,
			UserName: username,
			Year:     year,
		})
		return true
	}
	return false

}

//Annually Goal Function

func GetAnnuallyGIdByMonthlyGId(id int) int {
	mg, err := gorm.G[MonthlyGoal](db).Where("monthlyGId=?", id).Select("annuallyGId").Find(ctx)
	if err != nil {
		return -1
	}
	return mg[0].AGID
}

func GetAnnuallyGs(id int) ([]AnnuallyGoal, error) {
	ag, err := gorm.G[AnnuallyGoal](db).Where("annuallyPId=?", id).Find(ctx)
	if err != nil {
		return nil, err
	}
	return ag, err

}

func GetAnnuallyGById(annuallyGId int) (AnnuallyGoal, error) {
	ag, err := gorm.G[AnnuallyGoal](db).Where("annuallyGId=?", annuallyGId).Find(ctx)
	if err != nil {
		return AnnuallyGoal{AGID: annuallyGId}, err
	}
	return ag[0], err
}

func AddAnnuallyG(annuallyGoal *AnnuallyGoal) bool {
	gorm.G[AnnuallyGoal](db).Create(ctx, annuallyGoal)
	return true
}

func UpdateAnnuallyG(annuallyGoal *AnnuallyGoal) bool {
	gorm.G[AnnuallyGoal](db).Where("annuallyGId=?", annuallyGoal.AGID).Updates(ctx, *annuallyGoal)
	return true
}

func DeleteAnnuallyG(annuallyGId int) bool {
	gorm.G[AnnuallyGoal](db).Where("annuallyGId=?", annuallyGId).Delete(ctx)
	return true
}

//Annually Report

func GetAProgresses(annuallyPs []*AnnuallyPlan) (map[int]int, error) {
	var progresses = make(map[int]int)
	for _, annuallyP := range annuallyPs {
		ag, err := gorm.G[AnnuallyGoal](db).Where("annuallyPId=?", annuallyP.APID).Select("progress").Find(ctx)
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
		progresses[annuallyP.Year] = progress
	}
	return progresses, nil
}
