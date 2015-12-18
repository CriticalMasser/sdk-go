package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Event struct {
	APIKey     string `json:"api_key"`
	APISecret  string `json:"api_secret"`
	EventType  string `json:"event_type"`
	FacebookID int64  `json:"facebook_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Gender     string `json:"gender"`
	Email      string `json:"email"`
}

var (
	baseURL  = "https://a.criticalmasser.com/push/"
	jsonType = "application/json"
)

func SendEvent(e Event) error {
	body, err := json.Marshal(e)
	resp, err := http.Post(baseURL, jsonType, bytes.NewReader(body))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		body, err = ioutil.ReadAll(resp.Body)
		fmt.Printf("RB: %s", body)
	}
	return nil
}

func main() {
	e := Event{
		APIKey:     "1234567890",
		APISecret:  "1234567890",
		EventType:  "login",
		FacebookID: 12342352,
		FirstName:  "John",
		LastName:   "Doe",
		Gender:     "male",
		Email:      "john@doe.com",
	}

	if err := SendEvent(e); err != nil {
		log.Fatal(err)
	}
}
