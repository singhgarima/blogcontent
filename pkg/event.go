package pkg

import (
	"encoding/json"
)

type Event struct {
	Title      string     `json:"title"`
	Conference Conference `json:"conference"`
	VideoLink  string     `json:"videoLink"`
}

func GenerateEventsFromJson(data string) (events []Event, err error) {
	err = json.Unmarshal([]byte(data), &events)
	return events, err
}
