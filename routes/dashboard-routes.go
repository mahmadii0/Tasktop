package routes

import (
	"Tasktop/controllers"

	"github.com/gorilla/mux"
)

var DashRegister = func(r *mux.Router) {
	r.HandleFunc("/dashboard", controllers.DashHandler)

	//Read Goals
	r.HandleFunc("/dashboard/daily-goals/{goalId:[0-99]+}", controllers.DailyPlan).Methods("GET")
	r.HandleFunc("/dashboard/monthly-goals/{goalId:[0-99]+}", controllers.MonthlyPlan).Methods("GET")
	r.HandleFunc("/dashboard/annually-goals/{goalId:[0-99]+}", controllers.AnnuallyGoals).Methods("GET")

	//Create Goals
	r.HandleFunc("/dashboard/daily-goals/create", controllers.CDailyGoal).Methods("POST")
	r.HandleFunc("/dashboard/monthly-goals/create", controllers.CMonthlyGoal).Methods("POST")
	r.HandleFunc("/dashboard/annually-goals/create", controllers.CAnnuallyGoal).Methods("POST")

	//Update Goals
	r.HandleFunc("/dashboard/daily-goals/{goalId:[0-99]+}", controllers.DailyPlan).Methods("PUT")
	r.HandleFunc("/dashboard/monthly-goals/{goalId:[0-99]+}", controllers.MonthlyPlan).Methods("PUT")
	r.HandleFunc("/dashboard/annually-goals/{goalId:[0-99]+}", controllers.AnnuallyGoals).Methods("PUT")

	//Delete Goals
	r.HandleFunc("/dashboard/daily-goals/{goalId:[0-99]+}", controllers.DailyPlan).Methods("DELETE")
	r.HandleFunc("/dashboard/monthly-goals/{goalId:[0-99]+}", controllers.MonthlyPlan).Methods("DELETE")
	r.HandleFunc("/dashboard/annually-goals/{goalId:[0-99]+}", controllers.AnnuallyGoals).Methods("DELETE")

	//Report
	r.HandleFunc("/dashboard/report", controllers.Report)
}
