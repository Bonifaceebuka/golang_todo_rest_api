package routes

import (
	controllers "github.com/Bonifaceebuka/golang_todo_rest_api/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/task/store", controllers.StoreTodo).Methods("POST", "OPTIONS")

	return router
}
