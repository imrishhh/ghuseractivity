package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/imrishhh/ghuseractivity/model"
	"github.com/imrishhh/ghuseractivity/parse"
	"github.com/joho/godotenv"
)

func main() {
	cmdArgs := os.Args
	if len(cmdArgs) < 2 {
		fmt.Println("Command Usage: ghuseractivity <username> <page>(Default=1)")
		fmt.Println()
		fmt.Println("Example:", "ghuseractivity torvalds 1")
		return
	}
	username := cmdArgs[1]
	var (
		page int
		err  error
	)
	if len(cmdArgs) > 2 {
		page, err = strconv.Atoi(cmdArgs[2])
		if err != nil || page < 1 {
			fmt.Println("Invalid Page Number using: 1")
			page = 1
		}
	}
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load secret env file")
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("Please register GITHUB_TOKEN in env file!")
	}
	token = strings.TrimSpace(token)

	req := prepareRequest(username, page, token)

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
	if len(events) == 0 {
		fmt.Printf("No activity found at page no. %d", page)
	}
	//
	// data, _ := json.MarshalIndent(events, "", " ")
	// fmt.Printf("%+v\n", string(data))
	for _, e := range events {
		printEventDetail(e)
	}
}

func prepareRequest(username string, page int, token string) *http.Request {
	reqURL := fmt.Sprintf("https://api.github.com/users/%s/events?page=%d", username, page)

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

func timeAgo(t time.Time) string {
	d := time.Since(t)
	switch {
	case d < time.Minute:
		return fmt.Sprintf("%ds ago", int(d.Seconds()))
	case d < time.Hour:
		return fmt.Sprintf("%dm ago", int(d.Minutes()))
	case d < 24*time.Hour:
		return fmt.Sprintf("%dh ago", int(d.Hours()))
	case d < 30*24*time.Hour:
		return fmt.Sprintf("%dd ago", int(d.Hours()/24))
	case d < 12*30*24*time.Hour:
		return fmt.Sprintf("%dmo ago", int(d.Hours()/(24*30)))
	default:
		return fmt.Sprintf("%dy ago", int(d.Hours()/(365*24)))
	}
}

func printEventDetail(event model.Event) {
	payload, err := parse.ParsePayLoad(event)
	if err != nil {
		return
	}
	eventTime, err := time.Parse(time.RFC3339, event.CreatedAt)
	if err != nil {
		fmt.Println("Invalid event time")
		return
	}
	ago := timeAgo(eventTime)
	switch payload.(type) {
	case model.PushEvent:
		fmt.Printf("pushed changes to %s - %s\n", event.Repo.Name, ago)
	case model.WatchEvent:
		fmt.Printf("starred repository %s - %s\n", event.Repo.Name, ago)
	case model.ForkEvent:
		fmt.Printf("forked a repository %s - %s\n", event.Repo.Name, ago)
	}
}
