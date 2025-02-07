package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Event GitHub APIから取得するイベント情報のうち、必要な情報を定義
type Event struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	CreatedAt string `json:"created_at"`
}

// fetchEvents 指定したユーザーのイベント情報を取得して、Eventのスライスとして返す
func fetchEvents(username string) ([]Event, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "github-activity-cli")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var events []Event
	if err := json.Unmarshal(body, &events); err != nil {
		return nil, err
	}
	return events, nil
}

// printEvents 取得したイベント情報を、JSTに変換して出力
func printEvents(events []Event) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Printf("failed to load JST timezone: %v", err)
		jst = time.FixedZone("JST", 9*60*60)
	}

	for _, e := range events {
		t, err := time.Parse(time.RFC3339, e.CreatedAt)
		if err != nil {
			log.Printf("Time parse error for event %s: %v", e.Type, err)
			continue
		}
		fmt.Printf("Event: %s, Repo: %s, Created At: %s\n", e.Type, e.Repo.Name, t.In(jst).Format(time.RFC3339))
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activity <username>")
		os.Exit(1)
	}
	username := os.Args[1]

	events, err := fetchEvents(username)
	if err != nil {
		log.Fatalf("Failed to fetch events: %v", err)
	}
	printEvents(events)
}
