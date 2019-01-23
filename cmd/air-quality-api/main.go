package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nlopes/slack"
	"gitlab.com/troylelandshields/airqualitygovernor/cmd/air-quality-api/handlers"
)

var (
	slackToken = ""
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

	slackClientID = os.Getenv("SLACK_CLIENT_ID")
	slackClientSecret = os.Getenv("SLACK_CLIENT_SECRET")

	api := slack.New(slackToken)

	handlers := handlers.New()

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