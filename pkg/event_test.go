package pkg

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
			args:    args{`[{"title":"test","conference":{"name":"ctest","date":"2021-01-01"},"videoLink":"http://test"}]`},
			wantLen: 1,
			want: []Event{
				{
					Title: "test",
					Conference: Conference{
						Name: "ctest",
						Date: "2021-01-01",
					},
					VideoLink: "http://test",
				},
			},
		},
		{
			name:    "Should return 2 events",
			args:    args{`[{"title":"test","conference":{"name":"ctest","date":"2021-01-01"},"videoLink":"http://test"},{"title":"test2","conference":{"name":"ctest2","date":"2022-01-01"},"videoLink":""}]`},
			wantLen: 2,
			want: []Event{
				{
					Title: "test",
					Conference: Conference{
						Name: "ctest",
						Date: "2021-01-01",
					},
					VideoLink: "http://test",
				}, {
					Title: "test2",
					Conference: Conference{
						Name: "ctest2",
						Date: "2022-01-01",
					},
					VideoLink: "",
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
