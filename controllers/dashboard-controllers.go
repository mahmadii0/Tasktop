package controllers

import (
	"net/http"
)

func DashHandler(w http.ResponseWriter, r *http.Request) {
	TemplateRender(w, "/dashboard/dashboard", nil)
}

//Read Goals

func DailyGoal(w http.ResponseWriter, r *http.Request) {

}
func MonthlyGoal(w http.ResponseWriter, r *http.Request) {

}
func AnnuallyGoal(w http.ResponseWriter, r *http.Request) {

}

//Create Goals

func CDailyGoal(w http.ResponseWriter, r *http.Request) {
	TemplateRender(w, "/dashboard/create-goals/daily/create-update", nil)
}

func CMonthlyGoal(w http.ResponseWriter, r *http.Request) {
	TemplateRender(w, "/dashboard/create-goals/monthly/create-update", nil)
}
func CAnnuallyGoal(w http.ResponseWriter, r *http.Request) {
	// if r.Method == "GET" {
	// 	TemplateRender(w, "/dashboard/create-goals/annually/create-update", nil)
	// } else if r.Method == "POST" {
	// 	if err := r.ParseForm(); err != nil {
	// 		log.Println(err)
	// 		return
	// 	}
	// 	title := r.FormValue("title")
	// 	desc := r.FormValue("description")
	// 	annuallyPId, err := models.GetAnnuallyPByUserId(1)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	annuallyG := &models.AnnuallyGoal{Title: title, Desc: desc, Status: true, APID: annuallyPId.APID}
	// 	models.AddAnnuallyG(annuallyG)
	// 	r.Method = "GET"
	// 	AnnuallyGoals(w, r)
	// } else {
	// 	fmt.Println("dd")
	// }
}

//Update Goals

func UDailyGoal(w http.ResponseWriter, r *http.Request) {

}

func UMonthlyGoal(w http.ResponseWriter, r *http.Request) {

}

func UAnnuallyGoals(w http.ResponseWriter, r *http.Request) {

}

//Delete Goals

func DDailyGoal(w http.ResponseWriter, r *http.Request) {

}

func DMonthlyGoal(w http.ResponseWriter, r *http.Request) {

}

func DAnnuallyGoal(w http.ResponseWriter, r *http.Request) {

}

//Report

func Report(w http.ResponseWriter, r *http.Request) {

}
