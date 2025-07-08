package routes

import (
	"Tasktop/controllers"

	"github.com/gorilla/mux"
)

var DashRegister = func(r *mux.Router) {
	r.HandleFunc("", controllers.DashHandler)
	//Read Goals
	r.HandleFunc("/daily-goals/{goalId:[0-99]+}", controllers.DailyGoal).Methods("GET")
	r.HandleFunc("/monthly-goals/{goalId:[0-99]+}", controllers.MonthlyGoal).Methods("GET")
	r.HandleFunc("/annually-goals/{goalId:[0-99]+}", controllers.AnnuallyGoal).Methods("GET")

	//Create Goals
	r.HandleFunc("/daily-goals/create", controllers.CDailyGoal).Methods("POST")
	r.HandleFunc("/monthly-goals/create", controllers.CMonthlyGoal).Methods("POST")
	r.HandleFunc("/annually-goals/create", controllers.CAnnuallyGoal).Methods("POST")

	//Update Goals
	r.HandleFunc("/daily-goals/{goalId:[0-99]+}", controllers.UDailyGoal).Methods("PUT")
	r.HandleFunc("/monthly-goals/{goalId:[0-99]+}", controllers.UMonthlyGoal).Methods("PUT")
	r.HandleFunc("/annually-goals/{goalId:[0-99]+}", controllers.UAnnuallyGoals).Methods("PUT")

	//Delete Goals
	r.HandleFunc("/daily-goals/{goalId:[0-99]+}", controllers.DDailyGoal).Methods("DELETE")
	r.HandleFunc("/monthly-goals/{goalId:[0-99]+}", controllers.DMonthlyGoal).Methods("DELETE")
	r.HandleFunc("/annually-goals/{goalId:[0-99]+}", controllers.DAnnuallyGoal).Methods("DELETE")

	//Report
	r.HandleFunc("/report", controllers.Report)
}
