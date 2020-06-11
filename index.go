package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

var ctx context.Context
var client *mongo.Client
var collection *mongo.Collection
var episodes []bson.M

func getAllRooms(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(episodes)
}

func main() {
	connection()
	fmt.Println("hello")
	r := mux.NewRouter()
	r.HandleFunc("/Rooms", getAllRooms)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func connection() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/?connect=direct"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	tarkovDatabase := client.Database("TarkvoDb")
	roomsCollection := tarkovDatabase.Collection("Rooms")
	//_, err = roomsCollection.InsertOne(ctx, bson.D{
	//	{"Rate", "20"},
	//	{"Host", "Loh"},
	//	{"Description", "Da"},
	//	{"Players", "15"},
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	cursor, err := roomsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
}