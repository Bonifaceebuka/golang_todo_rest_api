package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/Bonifaceebuka/golang_todo_rest_api/routes"
)

func main() {
	router := router.Router()
	port := ":8080"
	fmt.Println("Starting the server on port %s", port)

	log.Fatal(http.ListenAndServe(port, router))

	fmt.Println("Server started on port %s", port)

}
