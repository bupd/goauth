package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Payload struct {
	Token        string `json:"token"`
	Repo         string `json:"repo"`
	Ref          string `json:"ref"`
	Workflow     string `json:"workflow"`
	RunID        string `json:"run_id"`
	Actor        string `json:"actor"`
	EventName    string `json:"event_name"`
	CommitSHA    string `json:"commit_sha"`
}

func main() {
	webhookURL := os.Getenv("WEBHOOK_URL") // ✅ webhook passed as env var
	token := os.Getenv("GITHUB_TOKEN")

	if webhookURL == "" || token == "" {
		fmt.Println("Missing WEBHOOK_URL or GITHUB_TOKEN")
		return
	}

	payload := Payload{
		Token:     token,
		Repo:      os.Getenv("GITHUB_REPOSITORY"),
		Ref:       os.Getenv("GITHUB_REF"),
		Workflow:  os.Getenv("GITHUB_WORKFLOW"),
		RunID:     os.Getenv("GITHUB_RUN_ID"),
		Actor:     os.Getenv("GITHUB_ACTOR"),
		EventName: os.Getenv("GITHUB_EVENT_NAME"),
		CommitSHA: os.Getenv("GITHUB_SHA"),
	}

	data, _ := json.Marshal(payload)

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Sent payload, status:", resp.Status)
}
