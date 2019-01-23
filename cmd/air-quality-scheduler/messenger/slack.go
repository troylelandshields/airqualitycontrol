package messenger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type slackMessage struct {
	Text string `json:"text"`
}

func Send(webhookURL string, message string) error {
	msg := slackMessage{
		Text: message,
	}

	msgBody, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewReader(msgBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status code when sending message: %d", resp.StatusCode)
	}

	return nil
}
