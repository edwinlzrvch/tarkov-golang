package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/edwinlzrvch/tarkov-golang/pkg/entity"
	"github.com/edwinlzrvch/tarkov-golang/pkg/room"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx context.Context
var client *mongo.Client
var collection *mongo.Collection
var userService *room.Service
var rooms []*entity.Room

func getAllRooms(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(rooms)
}

func main() {
	connection()

	fmt.Println("hello")
	r := mux.NewRouter()
	r.HandleFunc("/Rooms", getAllRooms)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)
	srv := &http.Server{
		Handler: handler,
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
	userRepo := room.NewMongoRepository(ctx, tarkovDatabase)
	userService = room.NewService(userRepo)
	rooms = userService.GetAllRooms()
}
