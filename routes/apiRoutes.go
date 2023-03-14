package routes

import (
	controllers "github.com/Bonifaceebuka/golang_todo_rest_api/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Home).Methods("GET")

	return router
}
