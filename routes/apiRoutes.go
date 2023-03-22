package routes

import (
	controllers "github.com/Bonifaceebuka/golang_todo_rest_api/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/task/store", controllers.StoreTodo).Methods("POST", "OPTIONS")
	router.HandleFunc("/tasks", controllers.GetAllTodos).Methods("GET", "OPTIONS")
	router.HandleFunc("/task/{task_id}", controllers.GetTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/task/{task_id}", controllers.UpdateTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/task/{task_id}", controllers.DeleteTask).Methods("DELETE", "OPTIONS")

	return router
}
