package models

//import (
//	"Tasktop/configure"
//
//	"gorm.io/gorm"
//)
//
//func CheckStatus(typee string, id int) bool {
//	Status := true
//	switch typee {
//	case "daily":
//		query := `UPDATE dailygoals SET status= IF(status=1,0,1) WHERE dailyGId=?`
//		_, err := db.Exec(query, id)
//		if err != nil {
//			Status = false
//		}
//		// query = `SELECT status FROM dailygoals WHERE dailyGId=?`
//		// row := db.QueryRow(query, id)
//		// err = row.Scan(&status)
//		// if err != nil {
//		// 	Status = false
//		// }
//
//	case "monthly":
//		query := `UPDATE monthlygoals SET status= IF(status=1,0,1),progress=IF(status=1,status,100) WHERE monthlyGId=?`
//		_, err := db.Exec(query, id)
//		if err != nil {
//			Status = false
//		}
//	case "annually":
//		query := `UPDATE annuallygoals SET status= IF(status=1,0,1),progress=IF(status=1,status,100) WHERE annuallyGId=?`
//		_, err := db.Exec(query, id)
//		if err != nil {
//			Status = false
//		}
//	}
//	return Status
//}
//
//func SetProgress(typee string, id int) bool {
//	Status := true
//	switch typee {
//	case "monthly":
//		monthlyGId := GetMonthlyGIdByDailyGId(id)
//		if monthlyGId == -1 {
//			Status = false
//		}
//		query := `SELECT status FROM dailygoals WHERE monthlyGId=?`
//		rows, err := db.Query(query, monthlyGId)
//		if err != nil {
//			Status = false
//		}
//		statuses := make([]float32, 0)
//		doneCounter := 0.00
//		counter := 0.00
//		for rows.Next() {
//			var status int
//			if err := rows.Scan(&status); err != nil {
//				Status = false
//			}
//			if status == 0 {
//				doneCounter++
//			}
//			statuses = append(statuses, float32(status))
//			counter++
//		}
//		result := (doneCounter / counter) * 100
//		query = `UPDATE monthlyGoals SET progress=? WHERE monthlyGId=?`
//		_, err = db.Exec(query, int(result), monthlyGId)
//		if err != nil {
//			Status = false
//		}
//
//	case "annually":
//		annuallyGId := GetAnnuallyGIdByMonthlyGId(id)
//		if annuallyGId == -1 {
//			Status = false
//		}
//		query := `SELECT status FROM monthlygoals WHERE annuallyGId=?`
//		rows, err := db.Query(query, annuallyGId)
//		if err != nil {
//			Status = false
//		}
//		statuses := make([]float32, 0)
//		activeCounter := 0.00
//		counter := 0.00
//		for rows.Next() {
//			var status int
//			if err := rows.Scan(&status); err != nil {
//				Status = false
//			}
//			if status == 0 {
//				activeCounter++
//			}
//			statuses = append(statuses, float32(status))
//			counter++
//		}
//		result := (activeCounter / counter) * 100
//		query = `UPDATE annuallyGoals SET progress=? WHERE annuallyGId=?`
//		_, err = db.Exec(query, int(result), annuallyGId)
//		if err != nil {
//			Status = false
//		}
//	}
//
//	return Status
//}
