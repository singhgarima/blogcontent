package pkg

import "testing"

func TestConference_GetFormattedDate(t *testing.T) {
	tests := []struct {
		name string
		date string
		want string
	}{
		{name: "Should return formatted date", date: "2021-02-10", want: "10 Feb 2021"},
		{name: "Should return empty string for invalid date", date: "invalid", want: ""},
		{name: "Should return empty string for empty date", date: "invalid", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conference{
				Name: "Star Wars Conference",
				Date: tt.date,
			}
			if got := c.GetFormattedDate(); got != tt.want {
				t.Errorf("GetFormattedDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
