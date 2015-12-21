// cm Critical Masser SDK for Go.
package cm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type EventData struct {
	EventType  string `json:"event_type"`
	FacebookID int64  `json:"facebook_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Gender     string `json:"gender"`
	Email      string `json:"email"`
}

// Event describes an event and its attributes.
type Event struct {
	APIKey    string    `json:"API_KEY"`
	APISecret string    `json:"API_SECRET"`
	Data      EventData `json:"data"`
	SDK       string    `json:"sdk"`
}

var (
	baseURL  = "https://a.criticalmasser.com/push/"
	jsonType = "application/json"
	retries  = 3

	// EventLogin login event
	EventLogin = "login"

	// EventRegister login event
	EventRegister = "register"

	// SexMale type for male
	SexMale = "male"

	// SexFemale type for females
	SexFemale = "female"

	// SexUndefined type for undefined sex
	SexUndefined = "undefined"
)

var (
	// AccessKey AWS Access Key
	APIKey string
	// SecretKey AWS Secret Key
	APISecret string
)

// SendEvent sends event to Critical Masser's Analytics server, returns error
// if something went wrong.
func SendEvent(e Event) error {
	e.SDK = "sdk-go v0.1.0"
	e.APIKey = APIKey
	e.APISecret = APISecret

	body, err := json.Marshal(e)
	var resp *http.Response
	for i := 0; i < retries; i++ {
		resp, err = http.Post(baseURL, jsonType, bytes.NewReader(body))
		if err == nil && resp.StatusCode == http.StatusOK {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(http.StatusText(resp.StatusCode))
	}

	// Check for errors in response body
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to read response body: %s", err)
	}
	var data map[string]interface{}
	if err = json.Unmarshal(body, &data); err != nil {
		return fmt.Errorf("Failed to parse respose body: %s", err)
	}

	code, ok := data["code"].(float64)
	if !ok {
		return errors.New("Failed to determine respose status code")
	}
	status := int(code)

	if status != http.StatusOK {
		message, ok := data["message"].(string)
		if !ok {
			return fmt.Errorf("Received %d, couldn't retrieve message", status)
		}
		return fmt.Errorf("Received %d: %s", status, message)
	}
	return nil
}
