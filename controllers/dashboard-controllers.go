package controllers

import (
	"Tasktop/models"
	"Tasktop/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
)

func DashHandler(w http.ResponseWriter, r *http.Request) {
	TemplateRender(w, "/dashboard/dashboard", nil)
}

//Read Goal(single)

func DailyGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id, err := strconv.Atoi(vars["goalId"])
	if err != nil {
		http.Error(w, "Error while parsing goalId", http.StatusBadRequest)
	}
	dailyGoal, err := models.GetDailyGById(Id)
	if err != nil {
		http.Error(w, "Error while fetching data", http.StatusNotFound)
		return
	}
	if dailyGoal.DGID == 0 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dailyGoal)
}

func MonthlyGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id, err := strconv.Atoi(vars["goalId"])
	if err != nil {
		http.Error(w, "Error while parsing goalId", http.StatusBadRequest)
	}
	monthlyGoal, err := models.GetMonthlyGById(Id)
	if err != nil {
		http.Error(w, "Error while fetching data", http.StatusNotFound)
		return
	}
	if monthlyGoal.MGID == 0 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(monthlyGoal)
}
func AnnuallyGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id, err := strconv.Atoi(vars["goalId"])
	if err != nil {
		http.Error(w, "Error while parsing goalId", http.StatusBadRequest)
	}
	annuallyGoal, err := models.GetAnnuallyGById(Id)
	if err != nil {
		http.Error(w, "Error while fetching data", http.StatusNotFound)
		return
	}
	if annuallyGoal.AGID == 0 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(annuallyGoal)
}

//Read Goals(user goals on declear date)

func GetDailyGoals(w http.ResponseWriter, r *http.Request) {
	date := r.FormValue("date")
	st, _ := r.Cookie("session_token")
	username := models.GetUsernameBySessionToken(st.Value)
	id := models.GetDailyPId(username, date)
	dailyGoals, err := models.GetDailyGs(id)
	if err != nil {
		http.Error(w, "Error while fetching data", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	for _, dailyGoal := range dailyGoals {
		if dailyGoal.DGID == 0 {
			continue
		}
		json.NewEncoder(w).Encode(dailyGoal)
	}
}

func GetMonthlyGoals(w http.ResponseWriter, r *http.Request) {
	date := r.FormValue("date")
	st, _ := r.Cookie("session_token")
	username := models.GetUsernameBySessionToken(st.Value)
	id := models.GetMonthlyPId(username, date)
	monthlyGoals, err := models.GetMonthlyGs(id)
	if err != nil {
		http.Error(w, "Error while fetching data", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	for _, monthlyGoal := range monthlyGoals {
		if monthlyGoal.MGID == 0 {
			continue
		}
		json.NewEncoder(w).Encode(monthlyGoal)
	}
}

func GetAnnuallyGoals(w http.ResponseWriter, r *http.Request) {
	y := r.FormValue("year")
	y = y[0:]
	year, err := strconv.Atoi(y)
	if err != nil {
		http.Error(w, "Error while parsing year", http.StatusBadRequest)
	}
	st, _ := r.Cookie("session_token")
	username := models.GetUsernameBySessionToken(st.Value)
	id := models.GetAnnuallyPId(username, year)
	annuallyGoals, err := models.GetAnnuallyGs(id)
	if err != nil {
		http.Error(w, "Error while fetching data", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	for _, annuallyGoal := range annuallyGoals {
		if annuallyGoal.AGID == 0 {
			continue
		}
		json.NewEncoder(w).Encode(annuallyGoal)
	}
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
	y := r.FormValue("year")
	y = y[1:]
	year, err := strconv.Atoi(y)
	if err != nil {
		http.Error(w, "Error while parsing year", http.StatusBadRequest)
		return
	}

	st, _ := r.Cookie("session_token")
	username := models.GetUsernameBySessionToken(st.Value)
	id := models.GetAnnuallyPId(username, year)
	if id == 0 {
		status := models.AddAnnuallyP(username, year)
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
	var monthlyGoal models.MonthlyGoal
	vars := mux.Vars(r)
	Id, err := strconv.ParseInt(vars["goalId"], 0, 0)
	if err != nil {
		http.Error(w, "Error while parsing goalId", http.StatusBadRequest)
	}
	monthlyGoal.MGID = int(Id)
	monthlyGoal.Title = r.FormValue("title")
	monthlyGoal.Desc = r.FormValue("desc")
	monthlyGoal.Priority = r.FormValue("priority")
	p, err := strconv.Atoi(r.FormValue("progress"))
	if err != nil {
		http.Error(w, "Error while parsing goalId", http.StatusBadRequest)
		return
	}
	monthlyGoal.Progress = p
	s := r.FormValue("status")
	if s == "1" {
		monthlyGoal.Status = true
	} else if s == "0" {
		monthlyGoal.Status = false
	} else {
		http.Error(w, "Not valid status data", http.StatusBadRequest)
		return
	}
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
	status := models.UpdateMonthlyG(&monthlyGoal)
	if !(status) {
		http.Error(w, "Error while adding daily goal", http.StatusNotModified)
		return
	}
	return
}

func UAnnuallyGoals(w http.ResponseWriter, r *http.Request) {
	var annuallyGoal models.AnnuallyGoal
	vars := mux.Vars(r)
	Id, err := strconv.ParseInt(vars["goalId"], 0, 0)
	if err != nil {
		http.Error(w, "Error while parsing goalId", http.StatusBadRequest)
	}
	annuallyGoal.AGID = int(Id)
	annuallyGoal.Title = r.FormValue("title")
	annuallyGoal.Desc = r.FormValue("desc")
	annuallyGoal.Priority = r.FormValue("priority")
	p, err := strconv.Atoi(r.FormValue("progress"))
	if err != nil {
		http.Error(w, "Error while parsing progress", http.StatusBadRequest)
		return
	}
	annuallyGoal.Progress = p
	s := r.FormValue("status")
	if s == "1" {
		annuallyGoal.Status = true
	} else if s == "0" {
		annuallyGoal.Status = false
	} else {
		http.Error(w, "Not valid status data", http.StatusBadRequest)
		return
	}
	y := r.FormValue("year")
	y = y[1:]
	year, err := strconv.Atoi(y)
	if err != nil {
		http.Error(w, "Error while parsing year", http.StatusBadRequest)
		return
	}
	st, _ := r.Cookie("session_token")
	username := models.GetUsernameBySessionToken(st.Value)
	id := models.GetAnnuallyPId(username, year)
	if id == 0 {
		status := models.AddAnnuallyP(username, year)
		if !(status) {
			http.Error(w, "Error while adding annually plan", http.StatusNotModified)
			return
		}
	}
	annuallyGoal.APID = id
	status := models.UpdateAnnuallyG(&annuallyGoal)
	if !(status) {
		http.Error(w, "Error while adding daily goal", http.StatusNotModified)
		return
	}
	return
}

//Delete Goals

func DDailyGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id, err := strconv.Atoi(vars["goalId"])
	if err != nil {
		http.Error(w, "Error while parsing goalId", http.StatusBadRequest)
		return
	}
	status := models.DeleteDailyG(Id)
	if !(status) {
		http.Error(w, "Error while deleting daily goal", http.StatusBadRequest)
		fmt.Printf("Error while deleting daily goal")
		return
	}
	return
}

func DMonthlyGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id, err := strconv.Atoi(vars["goalId"])
	if err != nil {
		http.Error(w, "Error while parsing goalId", http.StatusBadRequest)
		return
	}
	status := models.DeleteMonthlyG(Id)
	if !(status) {
		http.Error(w, "Error while deleting monthly goal", http.StatusBadRequest)
		fmt.Printf("Error while deleting monthly goal")
		return
	}
	return
}

func DAnnuallyGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id, err := strconv.Atoi(vars["goalId"])
	if err != nil {
		http.Error(w, "Error while parsing goalId", http.StatusBadRequest)
		return
	}
	status := models.DeleteAnnuallyG(Id)
	if !(status) {
		http.Error(w, "Error while deleting annually goal", http.StatusBadRequest)
		fmt.Printf("Error while deleting annually goal")
		return
	}
	return
}

//Notes

func Notes(w http.ResponseWriter, r *http.Request) {
	st, _ := r.Cookie("session_token")
	username := models.GetUsernameBySessionToken(st.Value)
	notes, err := models.GetNotes(username)
	if err != nil {
		http.Error(w, "Error while fetching data", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	for _, note := range notes {
		json.NewEncoder(w).Encode(note)
	}
}

func CNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	note.Title = r.FormValue("title")
	note.NoteText = r.FormValue("noteText")
	st, _ := r.Cookie("session_token")
	note.UserName = models.GetUsernameBySessionToken(st.Value)
	status := models.AddNote(&note)
	if !(status) {
		http.Error(w, "Error while adding note", http.StatusBadRequest)
		return
	}
	return
}

func DNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id, err := strconv.Atoi(vars["noteId"])
	if err != nil {
		http.Error(w, "Error while parsing Id", http.StatusBadRequest)
		return
	}
	status := models.DeleteNote(Id)
	if !(status) {
		http.Error(w, "Error while deleteing note", http.StatusBadRequest)
		return
	}
	return
}

func DNotes() {
	loc, _ := time.LoadLocation("Asia/Tehran")
	c := cron.New(cron.WithLocation(loc))
	_, err := c.AddFunc("0 0 * * *", func() {
		models.DeleteAllNotes()
	})
	if err != nil {
		fmt.Printf("Error on deleteing notes: ", err)
	}

	c.Start()

	select {}
}

//Report

func Report(w http.ResponseWriter, r *http.Request) {

}
