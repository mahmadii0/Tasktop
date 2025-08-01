package routes

import (
	"Tasktop/controllers"

	"github.com/gorilla/mux"
)

var DashRegister = func(r *mux.Router) {
	r.HandleFunc("", controllers.DashHandler)
	//Read Goal(single)
	r.HandleFunc("/daily-goals/{goalId:[0-9]+}", controllers.DailyGoal).Methods("GET")
	r.HandleFunc("/monthly-goals/{goalId:[0-9]+}", controllers.MonthlyGoal).Methods("GET")
	r.HandleFunc("/annually-goals/{goalId:[0-9]+}", controllers.AnnuallyGoal).Methods("GET")

	//Read Goals(user goals on declear date)
	r.HandleFunc("/daily-goals", controllers.GetDailyGoals).Methods("GET")
	r.HandleFunc("/monthly-goals", controllers.GetMonthlyGoals).Methods("GET")
	r.HandleFunc("/annually-goals", controllers.GetAnnuallyGoals).Methods("GET")

	//Create Goals
	r.HandleFunc("/daily-goals/create", controllers.CDailyGoal).Methods("POST")
	r.HandleFunc("/monthly-goals/create", controllers.CMonthlyGoal).Methods("POST")
	r.HandleFunc("/annually-goals/create", controllers.CAnnuallyGoal).Methods("POST")

	//Update Goals
	r.HandleFunc("/daily-goals/{goalId:[0-9]+}", controllers.UDailyGoal).Methods("PUT")
	r.HandleFunc("/monthly-goals/{goalId:[0-9]+}", controllers.UMonthlyGoal).Methods("PUT")
	r.HandleFunc("/annually-goals/{goalId:[0-9]+}", controllers.UAnnuallyGoals).Methods("PUT")

	//Delete Goals
	r.HandleFunc("/daily-goals/{goalId:[0-9]+}", controllers.DDailyGoal).Methods("DELETE")
	r.HandleFunc("/monthly-goals/{goalId:[0-9]+}", controllers.DMonthlyGoal).Methods("DELETE")
	r.HandleFunc("/annually-goals/{goalId:[0-9]+}", controllers.DAnnuallyGoal).Methods("DELETE")

	//Notes
	r.HandleFunc("/notes", controllers.Notes).Methods("GET")
	r.HandleFunc("/notes/create", controllers.CNote).Methods("POST")
	r.HandleFunc("/notes/{noteId:[0-9]}", controllers.DNote).Methods("DELETE")

	//Report
	r.HandleFunc("/report", controllers.Report)
	r.HandleFunc("/report/daily", controllers.DailyReport).Methods("POST")
	r.HandleFunc("/report/monthly", controllers.MonthlyReport).Methods("GET")
	r.HandleFunc("/report/annually", controllers.AnnuallyReport).Methods("GET")

}
