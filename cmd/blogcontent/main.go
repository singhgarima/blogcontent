package main

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/singhgarima/blogcontent/pkg"
	"github.com/singhgarima/blogcontent/pkg/database"
)

func main() {
	cfg, err := loadDynamoDBConfig()
	if err != nil {
		panic(err)
	}

	var db = database.NewDynamoDB(dynamodb.NewFromConfig(cfg))

	err = initialiseDatabase(db)
	if err != nil {
		panic(err)
	}

	err = loadEvents(db)
	if err != nil {
		panic(err)
	}
}

func loadDynamoDBConfig() (aws.Config, error) {
	region := os.Getenv("AWS_REGION")
	if region == "localhost" {
		return loadDynamoDBConfigForLocal()
	} else {
		return loadDynamoDBConfigForAWS()
	}
}

func loadDynamoDBConfigForAWS() (aws.Config, error) {
	defaultRegion := "ap-southeast-2"
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = defaultRegion
	}
	return config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
}

func loadDynamoDBConfigForLocal() (aws.Config, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == dynamodb.ServiceID && region == "localhost" {
			return aws.Endpoint{
				URL: "http://localhost:8000",
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	localCredsProvider := credentials.StaticCredentialsProvider{
		Value: aws.Credentials{
			AccessKeyID: "dummy", SecretAccessKey: "dummy", SessionToken: "dummy",
			Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
		},
	}
	return config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(localCredsProvider),
	)
}

func initialiseDatabase(db *database.DynamoDB) error {
	tables := []database.Table{
		database.EventTable,
	}
	for _, table := range tables {
		err := db.CreateTable(table.Name, table.CreateTableInput)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadEvents(db *database.DynamoDB) error {
	raw, err := os.ReadFile("./content/events/events.json")
	if err != nil {
		return err
	}

	events, err := pkg.GenerateEventsFromJson(string(raw))
	if err != nil {
		return err
	}

	for _, event := range events {
		input := database.EventTable.PutItemFunc(event)
		err = db.PutItem(input)
		if err != nil {
			return err
		}
	}
	return err
}
