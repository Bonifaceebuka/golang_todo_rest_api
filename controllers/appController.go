package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Bonifaceebuka/golang_todo_rest_api/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var collection *mongo.Collection

func init() {
	LoadEnv()
	createDBInstance()
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error: unable to load the .env file")
	}
}

func createDBInstance() *mongo.Collection {

	connectionString := os.Getenv("MONGO_URL")
	DB_NAME := os.Getenv("DB_NAME")
	COLLECTION_NAME := os.Getenv("COLLECTION_NAME")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	usersCollection := client.Database(DB_NAME).Collection(COLLECTION_NAME)

	fmt.Println("connected to the DB")

	return usersCollection

}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-METHODS", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	strB, _ := json.Marshal("API service is fully up!")
	w.Write(strB)
	fmt.Println(string(strB))
}

func StoreTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-METHODS", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	r.ParseForm()
	// r.ParseMultipartForm()
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	fmt.Println(r.FormValue("task"))

	todo.Task = r.FormValue("task")
	coll := createDBInstance()
	results, err := coll.InsertOne(context.TODO(), todo)

	fmt.Println(results)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	strB, _ := json.Marshal("Todo task added successfully!")
	w.Write(strB)
	fmt.Println(string(strB))
}
