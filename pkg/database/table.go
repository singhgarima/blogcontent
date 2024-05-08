package database

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/singhgarima/blogcontent/pkg"
)

type Table struct {
	Name             string
	CreateTableInput dynamodb.CreateTableInput
	PutItemFunc      func(interface{}) dynamodb.PutItemInput
}

var EventTableName = "events"
var EventTable = Table{
	Name: EventTableName,
	CreateTableInput: dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{{
			AttributeName: aws.String("title"),
			AttributeType: types.ScalarAttributeTypeS,
		}, {
			AttributeName: aws.String("conferenceName"),
			AttributeType: types.ScalarAttributeTypeS,
		}, {
			AttributeName: aws.String("conferenceDate"),
			AttributeType: types.ScalarAttributeTypeS,
		}, {
			AttributeName: aws.String("videoLink"),
			AttributeType: types.ScalarAttributeTypeS,
		}},
		KeySchema: []types.KeySchemaElement{{
			AttributeName: aws.String("title"),
			KeyType:       types.KeyTypeHash,
		}, {
			AttributeName: aws.String("conferenceDate"),
			KeyType:       types.KeyTypeRange,
		}},
		LocalSecondaryIndexes: []types.LocalSecondaryIndex{{
			IndexName: aws.String("conferenceNameIndex"),
			KeySchema: []types.KeySchemaElement{{
				AttributeName: aws.String("title"),
				KeyType:       types.KeyTypeHash,
			}, {
				AttributeName: aws.String("conferenceName"),
				KeyType:       types.KeyTypeRange,
			}},
			Projection: &types.Projection{
				ProjectionType: types.ProjectionTypeAll,
			}}, {
			IndexName: aws.String("videoLinkIndex"),
			KeySchema: []types.KeySchemaElement{{
				AttributeName: aws.String("title"),
				KeyType:       types.KeyTypeHash,
			}, {
				AttributeName: aws.String("videoLink"),
				KeyType:       types.KeyTypeRange,
			}},
			Projection: &types.Projection{
				ProjectionType: types.ProjectionTypeAll,
			},
		}},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(1),
		},
		TableName: aws.String(EventTableName),
	},
	PutItemFunc: func(item interface{}) dynamodb.PutItemInput {
		event := item.(pkg.Event)
		return dynamodb.PutItemInput{
			Item: map[string]types.AttributeValue{
				"title":          &types.AttributeValueMemberS{Value: event.Title},
				"conferenceName": &types.AttributeValueMemberS{Value: event.ConferenceName},
				"conferenceDate": &types.AttributeValueMemberS{Value: event.ConferenceDate},
				"videoLink":      &types.AttributeValueMemberS{Value: event.VideoLink},
			},
			TableName: aws.String(EventTableName),
		}
	},
}
