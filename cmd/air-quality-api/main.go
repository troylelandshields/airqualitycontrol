package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/troylelandshields/airqualitycontrol/cmd/air-quality-api/handlers"
	"github.com/troylelandshields/airqualitycontrol/webhooks"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("No port!")
		os.Exit(1)
	}

	connectionString := os.Getenv("DATABASE_URL")

	db, err := connectDatabase(connectionString)
	if err != nil {
		fmt.Println("Error connecting to database", err)
		os.Exit(1)
	}

	slackClientID := os.Getenv("SLACK_CLIENT_ID")
	slackClientSecret := os.Getenv("SLACK_CLIENT_SECRET")

	webhooksClient := webhooks.New(db)

	handlers := handlers.New(slackClientID, slackClientSecret, webhooksClient)

	router := mux.NewRouter()
	router.HandleFunc("/api/slack/redirect", handlers.AuthRedirect)

	fmt.Println("Waiting for requests on port:", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println("Done", err)
	}
}

func connectDatabase(connectionString string) (*sql.DB, error) {

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
