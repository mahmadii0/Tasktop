package routes

import (
	"Tasktop/controllers"
	"github.com/gorilla/mux"
)

var DashRegister = func(r *mux.Router) {
	r.HandleFunc("/dashboard", controllers.DashHandler)
	//Plans
	r.HandleFunc("/dashboard/daily-plans", controllers.DailyPlan)
	r.HandleFunc("/dashboard/monthly-plans", controllers.MonthlyPlan)
	r.HandleFunc("/dashboard/annually-plans", controllers.AnnuallyPlan)
	//Read
	r.HandleFunc("/dashboard/daily-goals", controllers.DailyPlan)
	r.HandleFunc("/dashboard/monthly-goals", controllers.MonthlyPlan)
	r.HandleFunc("/dashboard/annually-goals", controllers.AnnuallyGoals).Methods("GET")
	r.HandleFunc("/dashboard/annually-goals", controllers.AnnuallyGoals).Methods("POST")
	//Create
	r.HandleFunc("/dashboard/daily-goals/create", controllers.CDailyGoal).Methods("GET")
	r.HandleFunc("/dashboard/daily-goals/create", controllers.CDailyGoal).Methods("POST")
	r.HandleFunc("/dashboard/monthly-goals/create", controllers.CMonthlyGoal).Methods("GET")
	r.HandleFunc("/dashboard/monthly-goals/create", controllers.CMonthlyGoal).Methods("POST")
	r.HandleFunc("/dashboard/annually-goals/create", controllers.CAnnuallyGoal).Methods("GET")
	r.HandleFunc("/dashboard/annually-goals/create", controllers.CAnnuallyGoal).Methods("POST")

	//Update
	r.HandleFunc("/dashboard/annually-goals/update/{goalId:[0-99]+}", controllers.CAnnuallyGoal).Methods("GET")
	r.HandleFunc("/dashboard/annually-goals/update/{goalId:[0-99]+}", controllers.CAnnuallyGoal).Methods("POST")
	//Report
	r.HandleFunc("/dashboard/report", controllers.Report)
}
