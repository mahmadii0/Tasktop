package controllers

import (
	"Tasktop/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func DashHandler(w http.ResponseWriter, r *http.Request) {
	TemplateRender(w, "/dashboard/dashboard", nil)
}

func DailyPlan(w http.ResponseWriter, r *http.Request) {

}

func MonthlyPlan(w http.ResponseWriter, r *http.Request) {

}

func AnnuallyPlan(w http.ResponseWriter, r *http.Request) {
	TemplateRender(w, "/dashboard/annually-plans", nil)
}
func AnnuallyGoals(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		TemplateRender(w, "/dashboard/annually-goals", nil)
	} else if r.Method == "POSt" {
		if err := r.ParseForm(); err != nil {
			log.Println(err)
			return
		}
		title := r.FormValue("title")
		desc := r.FormValue("description")
		annaullyPId := r.FormValue("APID")
		conv, err := strconv.Atoi(annaullyPId)
		if err != nil {
			log.Println(err)
			return
		}
		annuallyG := &models.AnnuallyGoal{Title: title, Desc: desc, Status: true, APID: conv}
		models.AddAnnuallyG(annuallyG)
	}
}

//Create Controllers

func CDailyGoal(w http.ResponseWriter, r *http.Request) {
	TemplateRender(w, "/dashboard/create-goals/daily/create-update", nil)
}

func CMonthlyGoal(w http.ResponseWriter, r *http.Request) {
	TemplateRender(w, "/dashboard/create-goals/monthly/create-update", nil)
}
func CAnnuallyGoal(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		TemplateRender(w, "/dashboard/create-goals/annually/create-update", nil)
	} else if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			log.Println(err)
			return
		}
		title := r.FormValue("title")
		desc := r.FormValue("description")
		annuallyPId, err := models.GetAnnuallyPByUserId(1)
		if err != nil {
			log.Println(err)
		}
		annuallyG := &models.AnnuallyGoal{Title: title, Desc: desc, Status: true, APID: annuallyPId.APID}
		models.AddAnnuallyG(annuallyG)
		r.Method = "GET"
		AnnuallyGoals(w, r)
	} else {
		fmt.Println("dd")
	}
}

func Report(w http.ResponseWriter, r *http.Request) {

}
