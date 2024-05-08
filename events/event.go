package events

import (
	"encoding/json"
	"time"
)

type Event struct {
	Title          string `json:"title"`
	ConferenceName string `json:"conferenceName"`
	ConferenceDate string `json:"conferenceDate"`
	VideoLink      string `json:"videoLink"`
}

func (e *Event) GetFormattedDate() string {
	parseDate, err := time.Parse(time.DateOnly, e.ConferenceDate)
	if err != nil {
		return ""
	}
	return parseDate.Format("02 Jan 2006")
}

func GenerateEventsFromJson(data string) (events []Event, err error) {
	err = json.Unmarshal([]byte(data), &events)
	return events, err
}
