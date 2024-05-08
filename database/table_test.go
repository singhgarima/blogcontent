package database

import (
	"github.com/singhgarima/blogcontent/events"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func TestEventTable_PutItem(t *testing.T) {
	event := events.Event{
		Title:          "test",
		ConferenceName: "ctest",
		ConferenceDate: "2021-01-01",
		VideoLink:      "http://test",
	}

	input := EventTable.PutItemFunc(event)

	if *input.TableName != EventTable.Name {
		t.Errorf("TableName is not set correctly")
	}

	item := input.Item
	gotTitle := item["title"].(*types.AttributeValueMemberS).Value
	if gotTitle != event.Title {
		t.Errorf("title = %v, want %v", gotTitle, event.Title)
	}
}
