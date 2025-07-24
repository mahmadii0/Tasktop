package controllers

import (
	"Tasktop/models"
	"Tasktop/utils"
	"net/http"
	"strconv"
	"time"
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
	var dailyGoal models.DailyGoal
	dailyGoal.Title = r.FormValue("title")
	timeTD := r.FormValue("timeTD")
	parsedTime, _ := time.Parse(time.RFC3339, timeTD)
	dailyGoal.TimeTD = parsedTime
	dailyGoal.Priority = r.FormValue("priority")
	dailyPId, _ := strconv.Atoi(r.FormValue("dailyPId"))
	dailyGoal.DPID = dailyPId
	mGId := r.FormValue("monthlyGId")
	var monthlyGId int = 0
	if mGId != "" {
		monthlyGId, _ = strconv.Atoi(mGId)
	}
	dailyGoal.MGID = monthlyGId
	st, _ := r.Cookie("session_token")
	username := models.GetUsernameBySessionToken(st.Value)
	Id, err := utils.ConvertTimeToPlanId(timeTD)
	if err != nil {
		er := http.StatusNotModified
		http.Error(w, "Error while Id converting", er)
		return
	}
	status := models.DailyPExist(username, Id)
	if !(status) {
		status = models.AddDailyP(Id, username)
		if !(status) {
			http.Error(w, "Error while add daily plan", http.StatusNotModified)
			return
		}
	}

	status = models.AddDailyG(&dailyGoal)

}

func CMonthlyGoal(w http.ResponseWriter, r *http.Request) {

}
func CAnnuallyGoal(w http.ResponseWriter, r *http.Request) {

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
