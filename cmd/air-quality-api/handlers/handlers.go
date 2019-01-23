package handlers

import (
	"fmt"
	"net/http"

	"github.com/nlopes/slack"
	"github.com/troylelandshields/airqualitygovernor/webhooks"
)

type Handler struct {
	clientID       string
	clientSecret   string
	webhooksClient *webhooks.Client
}

func New(clientID string, clientSecret string, webhooksClient *webhooks.Client) *Handler {
	return &Handler{
		clientID:       clientID,
		clientSecret:   clientSecret,
		webhooksClient: webhooksClient,
	}
}

func (h *Handler) AuthRedirect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	code := r.URL.Query().Get("code")

	resp, err := slack.GetOAuthResponseContext(ctx, http.DefaultClient, h.clientID, h.clientSecret, code, "")
	if err != nil {
		fmt.Println("Error getting oauth token", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.webhooksClient.Create(ctx, resp.IncomingWebhook.URL)
	if err != nil {
		fmt.Println("error creating webhook", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
