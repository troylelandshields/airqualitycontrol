package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/troylelandshields/airqualitygovernor/cmd/air-quality-scheduler/airquaility"
	"github.com/troylelandshields/airqualitygovernor/cmd/air-quality-scheduler/messenger"
	"github.com/troylelandshields/airqualitygovernor/webhooks"
)

func main() {
	airNowAPIKey := os.Getenv("AIR_NOW_API_KEY")

	airQuality, err := airquaility.AirQuality("84094", airNowAPIKey)
	if err != nil {
		fmt.Println("error getting air quality", err)
		os.Exit(1)
	}

	if !airQuality.ShouldSend() {
		fmt.Println("Don't need to send a message")
		os.Exit(0)
	}

	connectionString := os.Getenv("DATABASE_URL")

	db, err := connectDatabase(connectionString)
	if err != nil {
		fmt.Println("Error connecting to database", err)
		os.Exit(1)
	}

	webhooksClient := webhooks.New(db)

	webhooks, err := webhooksClient.Webhooks(context.Background())
	if err != nil {
		fmt.Println("Error getting webhooks", err)
		os.Exit(1)
	}

	for _, w := range webhooks {
		err = messenger.Send(w, airQuality.Message())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func connectDatabase(connectionString string) (*sql.DB, error) {

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
