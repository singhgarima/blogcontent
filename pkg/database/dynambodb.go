package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/aws/aws-sdk-go-v2/aws"
)

const defaultRegion = "ap-southeast-2"

type DynamoDB struct {
	Client DynamoDBClient
}

func NewDynamoDB(client DynamoDBClient) *DynamoDB {
	return &DynamoDB{
		Client: client,
	}
}

func (a *DynamoDB) TableExists(tableName string) (bool, error) {
	exists := true
	_, err := a.Client.DescribeTable(
		context.TODO(), &dynamodb.DescribeTableInput{TableName: aws.String(tableName)},
	)
	if err != nil {
		var notFoundEx *types.ResourceNotFoundException
		if errors.As(err, &notFoundEx) {
			err = nil
		}
		exists = false
	}
	return exists, err
}

func (a *DynamoDB) DeleteTable(tableName string) error {
	_, err := a.Client.DeleteTable(context.TODO(), &dynamodb.DeleteTableInput{
		TableName: aws.String(tableName)})
	return err
}

func (a *DynamoDB) CreateTable(tableName string, input dynamodb.CreateTableInput) (err error) {
	exists, err := a.TableExists(tableName)
	fmt.Println(exists, err)
	if exists {
		return nil
	}
	_, err = a.Client.CreateTable(context.TODO(), &input)
	if err != nil {
		return err
	}
	return a.waitForTable(tableName)
}

func (a *DynamoDB) PutItem(input dynamodb.PutItemInput) (err error) {
	_, err = a.Client.PutItem(context.TODO(), &input)
	return err
}

func (a *DynamoDB) waitForTable(tn string) error {
	w := dynamodb.NewTableExistsWaiter(a.Client)
	err := w.Wait(context.TODO(),
		&dynamodb.DescribeTableInput{
			TableName: aws.String(tn),
		},
		2*time.Minute,
		func(o *dynamodb.TableExistsWaiterOptions) {
			o.MaxDelay = 5 * time.Second
			o.MinDelay = 5 * time.Second
		})
	if err != nil {
		log.Printf("Wait for table %v exists failed. Here's why: %v\n", tn, err)
	}
	return err
}
