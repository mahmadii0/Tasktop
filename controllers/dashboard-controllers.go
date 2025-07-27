package controllers

import (
	"Tasktop/models"
	"Tasktop/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	date, _ := utils.SeparateDateTime(timeTD)
	dailyGoal.TimeTD = timeTD[1:]
	dailyGoal.Priority = r.FormValue("priority")
	mGId := r.FormValue("monthlyGId")
	var monthlyGId int = 0
	if mGId != "" {
		monthlyGId, _ = strconv.Atoi(mGId)
	}
	dailyGoal.MGID = monthlyGId
	st, _ := r.Cookie("session_token")
	username := models.GetUsernameBySessionToken(st.Value)
	id := models.GetDailyPId(username, date)
	if id == 0 {
		status := models.AddDailyP(username, date)
		if !(status) {
			http.Error(w, "Error while adding daily plan", http.StatusNotModified)
			return
		}
	}
	dailyGoal.DPID = id
	status := models.AddDailyG(&dailyGoal)
	if !(status) {
		http.Error(w, "Error while adding daily goal", http.StatusNotModified)
		return
	}
	return
}

func CMonthlyGoal(w http.ResponseWriter, r *http.Request) {
	var monthlyGoal models.MonthlyGoal
	monthlyGoal.Title = r.FormValue("title")
	monthlyGoal.Desc = r.FormValue("desc")
	monthlyGoal.Priority = r.FormValue("priority")
	date := r.FormValue("date")

	aGId := r.FormValue("annuallyGId")
	var annuallyGId int = 0
	if aGId != "" {
		annuallyGId, _ = strconv.Atoi(aGId)
	}
	monthlyGoal.AGID = annuallyGId
	st, _ := r.Cookie("session_token")
	username := models.GetUsernameBySessionToken(st.Value)
	id := models.GetMonthlyPId(username, date)
	if id == 0 {
		status := models.AddMonthlyP(username, date)
		if !(status) {
			http.Error(w, "Error while adding daily plan", http.StatusNotModified)
			return
		}
	}
	monthlyGoal.MPID = id
	status := models.AddMonthlyG(&monthlyGoal)
	if !(status) {
		http.Error(w, "Error while adding daily goal", http.StatusNotModified)
		return
	}
	return
}
func CAnnuallyGoal(w http.ResponseWriter, r *http.Request) {
	var annuallyGoal models.AnnuallyGoal
	annuallyGoal.Title = r.FormValue("title")
	annuallyGoal.Desc = r.FormValue("desc")
	annuallyGoal.Priority = r.FormValue("priority")
	year := r.FormValue("year")

	st, _ := r.Cookie("session_token")
	username := models.GetUsernameBySessionToken(st.Value)
	id := models.GetAnnuallyPId(username, year)
	if id == 0 {
		status := models.AddMonthlyP(username, year)
		if !(status) {
			http.Error(w, "Error while adding daily plan", http.StatusNotModified)
			return
		}
	}
	annuallyGoal.APID = id
	status := models.AddAnnuallyG(&annuallyGoal)
	if !(status) {
		http.Error(w, "Error while adding daily goal", http.StatusNotModified)
		return
	}
	return
}

//Update Goals

func UDailyGoal(w http.ResponseWriter, r *http.Request) {
	var dailyGoal models.DailyGoal
	vars := mux.Vars(r)
	Id, err := strconv.ParseInt(vars["goalId"], 0, 0)
	if err != nil {
		http.Error(w, "Error while parsing goalId", http.StatusBadRequest)
	}
	dailyGoal.DGID = int(Id)
	dailyGoal.Title = r.FormValue("title")
	timeTD := r.FormValue("timeTD")
	date, _ := utils.SeparateDateTime(timeTD)
	dailyGoal.TimeTD = timeTD[1:]
	dailyGoal.Priority = r.FormValue("priority")
	s := r.FormValue("status")
	if s == "1" {
		dailyGoal.Status = true
	} else if s == "0" {
		dailyGoal.Status = false
	} else {
		http.Error(w, "Not valid status data", http.StatusBadRequest)
		return
	}
	mGId := r.FormValue("monthlyGId")
	var monthlyGId int = 0
	if mGId != "" {
		monthlyGId, _ = strconv.Atoi(mGId)
	}
	dailyGoal.MGID = monthlyGId
	st, _ := r.Cookie("session_token")
	username := models.GetUsernameBySessionToken(st.Value)
	id := models.GetDailyPId(username, date)
	if id == 0 {
		status := models.AddDailyP(username, date)
		if !(status) {
			http.Error(w, "Error while adding daily plan", http.StatusNotModified)
			return
		}
	}
	dailyGoal.DPID = id
	status := models.UpdateDailyG(&dailyGoal)
	if !(status) {
		http.Error(w, "Error while adding daily goal", http.StatusNotModified)
		return
	}
	return
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
