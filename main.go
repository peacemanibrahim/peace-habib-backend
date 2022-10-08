package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	db "github.com/peace-habib-exchange/backend/persistence"
	"github.com/peace-habib-exchange/backend/utility"
)

func main() {
	fmt.Println("Lailahailalah") // The Word that sustains the Heavens and the Earth.
	config, err := utility.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	database := db.Repository{}
	// Database Connection Below
	client, err := db.DatabaseConnection(config.DATABASE_URL)
	if err != nil {
		log.Fatal("Unable to connect to mongoDB databases")
	}
	databaseClient := database.NewDatabase(client, config.DATABASE_NAME)
	ctx := context.Background()
	if err := databaseClient.CreateIndex(ctx); err != nil {
		log.Fatal("Unable to create index")
	}

	router := mux.NewRouter()
	router.HandleFunc("/signup", CreateUser).Methods("POST")
	http.ListenAndServe(":7000", router)
}
