# Critical Masser SDK for Go
This SDK tries to make it easy for you to send data to Critical Masser, see the examples.

## Docs
Event describes an event and its attributes.

```go
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
```

SendEvent sends event to Critical Masser's Analytics server, returns error if something went wrong.

```go
func SendEvent(e Event) error
```

Both **APIKey** and **APISecret** will be provided by Critical Masser.

## Examples

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/criticalmasser/sdk-go"
)

func main() {
	cm.APIKey = os.Getenv("API_KEY")
	cm.APISecret = os.Getenv("API_SECRET")
	e := cm.Event{
		Data: cm.EventData{
			EventType:  cm.EventRegister,
			FacebookID: 12342352,
			FirstName:  "Rocky",
			LastName:   "Balboa",
			Gender:     cm.SexMale,
			Email:      "rocky@balboa.com",
		},
	}

	if err := cm.SendEvent(e); err != nil {
		log.Fatalf("Expected to succeed, but failed with error: %s", err)
	}
	fmt.Println("Done!")
}
```
