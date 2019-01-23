package main

import (
	"fmt"
	"os"

	"gitlab.com/troylelandshields/airqualitygovernor/cmd/air-quality-scheduler/airquaility"
	"gitlab.com/troylelandshields/airqualitygovernor/cmd/air-quality-scheduler/messenger"
	"gitlab.com/troylelandshields/airqualitygovernor/cmd/air-quality-scheduler/webhooks"
)

const (
	airNowAPIKey = "15EFF506-1A12-408C-AA60-AE4E95E83078"
)

func main() {
	airQuality, err := airquaility.AirQuality("84094", airNowAPIKey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if !airQuality.ShouldSend() {
		fmt.Println("Done")
		os.Exit(0)
	}

	webhooks := webhooks.Webhooks()

	for _, w := range webhooks {
		err = messenger.Send(w, airQuality.Message())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

}
