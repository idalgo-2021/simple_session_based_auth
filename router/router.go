package router

import (
	"simple_session_based_auth/controllers"
	"simple_session_based_auth/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.SessionMiddleware)

	router.HandleFunc("/login", controllers.LogIn).Methods("POST", "OPTIONS")
	router.HandleFunc("/refresh", controllers.Refresh).Methods("POST", "OPTIONS")
	router.HandleFunc("/logout", controllers.LogOut).Methods("POST", "OPTIONS")
	router.HandleFunc("/items", controllers.GetItems).Methods("GET", "OPTIONS")

	return router
}
