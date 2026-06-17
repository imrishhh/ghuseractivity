package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/imrishhh/ghuseractivity/model"
	"github.com/imrishhh/ghuseractivity/parse"
	"github.com/joho/godotenv"
)

func main() {
	cmdArgs := os.Args
	if len(cmdArgs) < 2 {
		fmt.Println("Command Usage: ghuseractivity <username>")
		fmt.Println()
		fmt.Println("Example:", "ghuseractivity torvalds")
		return
	}
	username := cmdArgs[1]

	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load secret env file")
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("Please register GITHUB_TOKEN in env file!")
	}
	token = strings.TrimSpace(token)

	req := prepareRequest(username, token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Something went wrong while fetching request\n", err)
		return
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("Something went wrong while trying to close model.body:", err)
		}
	}()
	if err := getStatusCodeError(resp.StatusCode); err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var events []model.Event
	if err := json.Unmarshal(body, &events); err != nil {
		log.Fatal("JSON decoding failed:", err)
	}

	data, _ := json.MarshalIndent(events, "", " ")
	fmt.Printf("%+v\n", string(data))
	for _, e := range events {
		printEventDetail(e)
	}
}

func prepareRequest(username string, token string) *http.Request {
	reqURL := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		log.Fatal("Something went wrong while preparing new http request:", err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)

	return req
}

func getStatusCodeError(code int) error {
	switch code {
	case http.StatusOK:
		return nil
	case http.StatusNotFound:
		return fmt.Errorf("no user with such name")
	default:
		return fmt.Errorf("something went wrong while fetching data")
	}
}

func printEventDetail(event model.Event) {
	payload, err := parse.ParsePayLoad(event)
	if err != nil {
		return
	}
	switch payload.(type) {
	case model.PushEvent:
		fmt.Printf("pushed changes to %s - %s", event.Repo.Name, event.CreatedAt)
	case model.WatchEvent:
		fmt.Printf("watched %s - %s", event.Repo.Name, event.CreatedAt)
	case model.ForkEvent:
		fmt.Printf("forked %s repository - %s", event.Repo.Name, event.CreatedAt)
	}
	fmt.Println()
}
