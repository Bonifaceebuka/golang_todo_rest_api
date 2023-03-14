package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func createDBInstance() {
	connectionString := os.Getenv("MONGO_URL")
	DB_NAME := os.Getenv("DB_NAME")
	COLLECTION_NAME := os.Getenv("COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to the DB")

	client.Database(DB_NAME).Collection(COLLECTION_NAME)
	fmt.Println("collection created")
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
