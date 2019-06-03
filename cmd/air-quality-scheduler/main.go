package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/troylelandshields/airqualitycontrol/cmd/air-quality-scheduler/airquaility"
	"github.com/troylelandshields/airqualitycontrol/cmd/air-quality-scheduler/messenger"
	"github.com/troylelandshields/airqualitycontrol/webhooks"
)

func main() {
	airNowAPIKey := os.Getenv("AIR_NOW_API_KEY")

	defaultTZ, err := time.LoadLocation("America/Denver")
	if err != nil {
		fmt.Println("Couldn't load default location")
		os.Exit(1)
	}
	t := time.Now().In(defaultTZ)

	switch t.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("enjoy the weekend")
		os.Exit(0)
	}

	airQuality, err := airquaility.AirQuality("84094", airNowAPIKey, t)
	if err != nil {
		fmt.Println("error getting air quality", err)
		os.Exit(1)
	}

	fmt.Println(airQuality.Message())

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
		msg := airQuality.Message()

		if w == "https://hooks.slack.com/services/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" {
			msg += " @triston.whetten"
		}

		err = messenger.Send(w, msg)
		if err != nil {
			fmt.Println("error sending to webhook", w, err)
			continue
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
