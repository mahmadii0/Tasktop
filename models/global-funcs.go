package models

import (
	"Tasktop/utils"

	"gorm.io/gorm"
)

func CheckStatus(typee string, id int) bool {
	Status := true
	switch typee {
	case "daily":
		db.Model(&DailyGoal{}).
			Where("dg_id = ?", id).
			Update("status", gorm.Expr("CASE WHEN status = 1 THEN 0 ELSE 1 END"))
	case "monthly":
		db.Model(&MonthlyGoal{}).
			Where("mg_id = ?", id).
			Update("status", gorm.Expr("CASE WHEN status = 1 THEN 0 ELSE 1 END"))
	case "annually":
		db.Model(&AnnuallyGoal{}).
			Where("ag_id = ?", id).
			Update("status", gorm.Expr("CASE WHEN status = 1 THEN 0 ELSE 1 END"))
	}
	return Status
}

func SetProgress(typee string, id int) bool {
	Status := true
	switch typee {
	case "monthly":
		monthlyGId := GetMonthlyGIdByDailyGId(id)
		if monthlyGId == -1 {
			Status = false
		}
		dg, err := gorm.G[DailyGoal](db).Where("mg_id = ?", id).Select("status").Find(ctx)
		if err != nil {
			Status = false
		}
		statuses := make([]float32, 0)
		doneCounter := 0.00
		counter := 0.00
		for _, item := range dg {
			status := utils.BoolToInt(item.Status)
			if status == 0 {
				doneCounter++
			}
			statuses = append(statuses, float32(status))
			counter++
		}
		result := (doneCounter / counter) * 100
		gorm.G[MonthlyGoal](db).Where("mg_id=?", id).Update(ctx, "progress", result)

	case "annually":
		annuallyGId := GetAnnuallyGIdByMonthlyGId(id)
		if annuallyGId == -1 {
			Status = false
		}
		mg, err := gorm.G[MonthlyGoal](db).Where("ag_id = ?", id).Select("status").Find(ctx)
		if err != nil {
			Status = false
		}
		statuses := make([]float32, 0)
		doneCounter := 0.00
		counter := 0.00
		for _, item := range mg {
			status := utils.BoolToInt(item.Status)
			if status == 0 {
				doneCounter++
			}
			statuses = append(statuses, float32(status))
			counter++
		}
		result := (doneCounter / counter) * 100
		gorm.G[AnnuallyGoal](db).Where("ag_id=?", id).Update(ctx, "progress", result)
	}
	return Status
}
