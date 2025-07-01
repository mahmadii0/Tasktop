package controllers

import (
	"Tasktop/models"
	"encoding/json"
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
	w.Header().Set("Content-Type", "application/json")
	var dailyGoal models.DailyGoal
	_ = json.NewDecoder(r.Body).Decode(&dailyGoal)
	status := models.AddDailyG(&dailyGoal)
	if status {
		json.NewEncoder(w).Encode(dailyGoal)
		DashHandler(w, r)
	}
	DashHandler(w, r)
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
