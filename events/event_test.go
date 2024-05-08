package events

import (
	"reflect"
	"testing"
)

func TestGenerateEventsFromJson(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
		want    []Event
		wantErr bool
	}{
		{
			name:    "Should return an event",
			args:    args{`[{"title":"test","conferenceName":"ctest","conferenceDate":"2021-01-01","videoLink":"http://test"}]`},
			wantLen: 1,
			want: []Event{
				{
					Title:          "test",
					ConferenceName: "ctest",
					ConferenceDate: "2021-01-01",
					VideoLink:      "http://test",
				},
			},
		},
		{
			name:    "Should return 2 events",
			args:    args{`[{"title":"test","conferenceName":"ctest","conferenceDate":"2021-01-01","videoLink":"http://test"},{"title":"test2","conferenceName":"ctest2","conferenceDate":"2022-01-01","videoLink":""}]`},
			wantLen: 2,
			want: []Event{
				{
					Title:          "test",
					ConferenceName: "ctest",
					ConferenceDate: "2021-01-01",
					VideoLink:      "http://test",
				}, {
					Title:          "test2",
					ConferenceName: "ctest2",
					ConferenceDate: "2022-01-01",
					VideoLink:      "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateEventsFromJson(tt.args.data)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewEventFromJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != tt.wantLen {
				t.Errorf("NewEventFromJson() got = %v, want %v", len(got), tt.wantLen)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventFromJson() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFormattedDate(t *testing.T) {
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
			e := &Event{
				Title:          "test",
				ConferenceName: "Star Wars Conference",
				ConferenceDate: tt.date,
				VideoLink:      "http://test",
			}
			if got := e.GetFormattedDate(); got != tt.want {
				t.Errorf("GetFormattedDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
