package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/imrishhh/ghuseractivity/parse"
	"github.com/imrishhh/ghuseractivity/response"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load secret env file")
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("Please register GITHUB_TOKEN in env file!")
	}

	token = strings.TrimSpace(token)

	req, err := http.NewRequest("GET", "https://api.github.com/users/torvalds/events", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal("Something went wrong! Events couldn't be fetched")
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("Something went wrong while trying to close response body:", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var events []response.Event
	if err := json.Unmarshal(body, &events); err != nil {
		log.Fatal("JSON decoding failed:", err)
	}
	data, _ := json.MarshalIndent(events, "", " ")
	fmt.Printf("%+v\n", string(data))

	for _, e := range events {
		printEventDetail(e)
	}
}

func printEventDetail(event response.Event) {
	payload, err := parse.ParsePayLoad(event)
	if err != nil {
		return
	}
	switch p := payload.(type) {
	case response.PushEvent:
		fmt.Printf("Pushed at %s", event.Repo.URL)
	case response.WatchEvent:
		fmt.Printf("Watched %s with action: %s", event.Repo.Name, p.Action)
	case response.ForkEvent:
		fmt.Printf("Forked: %s with action: %s", event.Repo.Name, p.Action)
	}
	fmt.Println()
}
