package routes

import (
	"Tasktop/controllers"

	"github.com/gorilla/mux"
)

var DashRegister = func(r *mux.Router) {
	r.HandleFunc("/dashboard", controllers.DashHandler)

	//Read Goals
	r.HandleFunc("/dashboard/daily-goals/{goalId:[0-99]+}", controllers.DailyGoal).Methods("GET")
	r.HandleFunc("/dashboard/monthly-goals/{goalId:[0-99]+}", controllers.MonthlyGoal).Methods("GET")
	r.HandleFunc("/dashboard/annually-goals/{goalId:[0-99]+}", controllers.AnnuallyGoal).Methods("GET")

	//Create Goals
	r.HandleFunc("/dashboard/daily-goals/create", controllers.CDailyGoal).Methods("POST")
	r.HandleFunc("/dashboard/monthly-goals/create", controllers.CMonthlyGoal).Methods("POST")
	r.HandleFunc("/dashboard/annually-goals/create", controllers.CAnnuallyGoal).Methods("POST")

	//Update Goals
	r.HandleFunc("/dashboard/daily-goals/{goalId:[0-99]+}", controllers.UDailyGoal).Methods("PUT")
	r.HandleFunc("/dashboard/monthly-goals/{goalId:[0-99]+}", controllers.UMonthlyGoal).Methods("PUT")
	r.HandleFunc("/dashboard/annually-goals/{goalId:[0-99]+}", controllers.UAnnuallyGoals).Methods("PUT")

	//Delete Goals
	r.HandleFunc("/dashboard/daily-goals/{goalId:[0-99]+}", controllers.DDailyGoal).Methods("DELETE")
	r.HandleFunc("/dashboard/monthly-goals/{goalId:[0-99]+}", controllers.DMonthlyGoal).Methods("DELETE")
	r.HandleFunc("/dashboard/annually-goals/{goalId:[0-99]+}", controllers.DAnnuallyGoal).Methods("DELETE")

	//Report
	r.HandleFunc("/dashboard/report", controllers.Report)
}
