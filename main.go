package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/peace-habib-exchange/backend/api"
	db "github.com/peace-habib-exchange/backend/persistence"
	"github.com/peace-habib-exchange/backend/utility"
)

var (
	port string
)

func main() {
	fmt.Println("Lailahailalah") // The Word that sustains the Heavens and the Earth.
	config, err := utility.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	database := db.PeaceRepository{}
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

	// Get the port number
	port = config.PORT

	router := mux.NewRouter()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPatch, http.MethodOptions})
	apiRouter := api.PeaceApi(router, databaseClient)
	cors := handlers.CORS(headersOk, methodsOk, originsOk)
	loggedRouter := handlers.LoggingHandler(os.Stdout, handlers.CompressHandler(cors(apiRouter)))

	// configure server
	server := &http.Server{
		Addr:           ":" + port,
		Handler:        loggedRouter,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()

	// router := mux.NewRouter()
	// router.HandleFunc("/signup", CreateUser).Methods("POST")
	// http.ListenAndServe(":7000", router)
}
