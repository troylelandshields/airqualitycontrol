package handlers

import (
	"fmt"
	"net/http"

	"github.com/nlopes/slack"
)

const ()

type Handler struct {
	clientID     string
	clientSecret string
}

func New(clientID string, clientSecret string, redirectURI string) *Handler {
	return &Handler{
		clientID:     clientID,
		clientSecret: clientSecret,
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

}
