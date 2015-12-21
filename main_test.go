package cm

import (
	"os"
	"testing"
)

func TestSendEvent(t *testing.T) {
	APIKey = os.Getenv("API_KEY")
	APISecret = os.Getenv("API_SECRET")
	e := Event{
		Data: EventData{
			EventType:  EventLogin,
			FacebookID: uint64(12342352),
			FirstName:  "Rocky",
			LastName:   "Balboa",
			Gender:     SexMale,
			Email:      "rocky@balboa.com",
		},
	}

	if err := SendEvent(e); err != nil {
		t.Fatalf("Expected to succeed, but failed with error: %s", err)
	}
}
