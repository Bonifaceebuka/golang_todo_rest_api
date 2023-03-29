package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Bonifaceebuka/golang_todo_rest_api/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var collection *mongo.Collection
var TaskModel []models.Todo

func init() {
	LoadEnv()
	collection = createDBInstance()
}

func LoadEnv() {
	// err := godotenv.Load(".env")
	err := godotenv.Load()
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
	resrponse, _ := json.Marshal("API service is fully up!")
	w.Write(resrponse)
	fmt.Println(string(resrponse))
}

func StoreTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	r.ParseForm()
	// r.ParseMultipartForm()
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	todo.Task = r.FormValue("task")
	results, err := collection.InsertOne(context.TODO(), todo)

	fmt.Println(results)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	resrponse, _ := json.Marshal("Todo task added successfully!")
	w.Write(resrponse)
	fmt.Println(string(resrponse))
}

func GetAllTodos(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET")

	tasks := getTodos()
	json.NewEncoder(res).Encode(tasks)
}

func getTodos() []primitive.M {
	data, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal("Unable to fetch the tasks")
	}

	var allTasks []primitive.M
	for data.Next(context.Background()) {
		var result bson.M
		err := data.Decode(&result)

		if err != nil {
			log.Fatal("Unable to decode this data")
		}

		allTasks = append(allTasks, result)

	}
	data.Close(context.Background())
	return allTasks
}

func GetTask(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Methods", "GET")
	res.Header().Set("Access-Control-Allow-Orign", "*")
	res.Header().Set("Access-Control-Allow-Type", "application/json")

	vars := mux.Vars(req)
	task_id := vars["task_id"]

	id, _ := primitive.ObjectIDFromHex(task_id)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	foundTask, err := collection.FindOne(ctx, bson.M{"_id": id}).DecodeBytes()

	defer cancel()
	if err != nil {
		log.Fatal("Unable to fetch the data for task")
	}

	fmt.Println(foundTask)
}

func UpdateTask(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Methods", "PUT")
	res.Header().Set("Access-Control-Allow-Orign", "*")
	res.Header().Set("Access-Control-Allow-Type", "application/json")

	vars := mux.Vars(req)
	task_id := vars["task_id"]

	id, _ := primitive.ObjectIDFromHex(task_id)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	foundTask, err := collection.FindOne(ctx, bson.M{"_id": id}).DecodeBytes()

	defer cancel()
	if err != nil {
		log.Fatal("Unable to fetch the data for task")
	}
	if foundTask == nil {
		log.Fatal("Task not found")
	}

	req.ParseForm()
	task := req.FormValue("task")
	status := req.FormValue("status")

	updatedTask, err := collection.UpdateOne(ctx, bson.M{"_id": id},
		bson.M{"$set": bson.M{"status": status, "task": task}})

	if err != nil {
		panic("Task was not updated")
	}

	fmt.Println(updatedTask)
}

func DeleteTask(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Methods", "DELETE")
	res.Header().Set("Access-Control-Allow-Orign", "*")
	res.Header().Set("Access-Control-Allow-Type", "application/json")

	vars := mux.Vars(req)
	task_id := vars["task_id"]

	id, _ := primitive.ObjectIDFromHex(task_id)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	foundTask, err := collection.FindOne(ctx, bson.M{"_id": id}).DecodeBytes()

	defer cancel()
	if err != nil {
		log.Fatal("Unable to fetch the data for task")
	}
	if foundTask == nil {
		log.Fatal("Task not found")
	}

	taskToDelete, err := collection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		panic("Task was not updated")
	}

	fmt.Println("Task deleted", taskToDelete)
}
