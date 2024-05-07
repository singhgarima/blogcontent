package pkg

import (
	"time"
)

type Conference struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

func (c *Conference) GetFormattedDate() string {
	parseDate, err := time.Parse(time.DateOnly, c.Date)
	if err != nil {
		return ""
	}
	return parseDate.Format("02 Jan 2006")
}
