package database

import (
	"errors"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/mock"
)

func TestNewDynamoDB(t *testing.T) {
	var mockClient DynamoDBClient = NewMockClient()

	var db DynamoDB
	NewDynamoDB(mockClient)

	if reflect.DeepEqual(&db.Client, &mockClient) {
		t.Errorf("Expected engine not set correctly, expected %s, but got %s", mockClient, db.Client)
	}
}

func TestDynamoDB_TableExists(t *testing.T) {
	tests := []struct {
		name               string
		stubDescribeOutput *dynamodb.DescribeTableOutput
		stubError          error
		wantExists         bool
		wantError          error
	}{
		{
			name:               "should return table exists",
			stubDescribeOutput: &dynamodb.DescribeTableOutput{},
			stubError:          nil,
			wantExists:         true,
			wantError:          nil,
		},
		{
			name:               "should return table does not exist",
			stubDescribeOutput: &dynamodb.DescribeTableOutput{},
			stubError:          &types.ResourceNotFoundException{},
			wantExists:         false,
			wantError:          nil,
		},
		{
			name:               "should return error",
			stubDescribeOutput: &dynamodb.DescribeTableOutput{},
			stubError:          errors.New("some error"),
			wantExists:         false,
			wantError:          errors.New("some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockClient = NewMockClient()
			var db DynamoDB
			db.Client = mockClient

			mockClient.On("DescribeTable", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.stubDescribeOutput, tt.stubError)

			exists, err := db.TableExists("test-table")

			mockClient.AssertNumberOfCalls(t, "DescribeTable", 1)

			if exists != tt.wantExists {
				t.Errorf("Expected table to exist, but it does not")
			}

			if !reflect.DeepEqual(err, tt.wantError) {
				t.Errorf("Expected error to be nil, but got %s", err)
			}
		})
	}

}

func TestDynamoDB_DeleteTable(t *testing.T) {
	tests := []struct {
		name       string
		stubOutput *dynamodb.DeleteTableOutput
		stubError  error
		wantError  error
	}{
		{
			name:       "should return nil for error",
			stubOutput: &dynamodb.DeleteTableOutput{},
			stubError:  nil,
			wantError:  nil,
		},
		{
			name:       "should return error",
			stubOutput: &dynamodb.DeleteTableOutput{},
			stubError:  errors.New("some error"),
			wantError:  errors.New("some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockClient = NewMockClient()
			var db DynamoDB
			db.Client = mockClient

			mockClient.On("DeleteTable", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.stubOutput, tt.stubError)

			err := db.DeleteTable("test-table")

			mockClient.AssertNumberOfCalls(t, "DeleteTable", 1)

			if !reflect.DeepEqual(err, tt.wantError) {
				t.Errorf("Expected error to be nil, but got %s", err)
			}
		})
	}
}

func TestDynamoDB_PutItem(t *testing.T) {
	tests := []struct {
		name          string
		putItemOutput *dynamodb.PutItemOutput
		stubError     error
		wantError     error
	}{
		{
			name:          "should return nil for error",
			putItemOutput: &dynamodb.PutItemOutput{},
			stubError:     nil,
			wantError:     nil,
		},
		{
			name:          "should return error",
			putItemOutput: &dynamodb.PutItemOutput{},
			stubError:     errors.New("some error"),
			wantError:     errors.New("some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockClient = NewMockClient()
			var db DynamoDB
			db.Client = mockClient

			mockClient.On("PutItem", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.putItemOutput, tt.stubError)

			err := db.PutItem(&dynamodb.PutItemInput{})

			mockClient.AssertNumberOfCalls(t, "PutItem", 1)

			if !reflect.DeepEqual(err, tt.wantError) {
				t.Errorf("Expected error to be nil, but got %s", err)
			}
		})
	}
}

func TestDynamoDB_CreateTable(t *testing.T) {
	tests := []struct {
		name              string
		stubOutput        *dynamodb.CreateTableOutput
		stubError         error
		stubDescribeError error
		wantError         error
	}{
		{
			name:              "should return nil for error and create a table",
			stubOutput:        &dynamodb.CreateTableOutput{},
			stubError:         nil,
			stubDescribeError: &types.ResourceNotFoundException{},
			wantError:         nil,
		},
		{
			name:              "should return error",
			stubOutput:        &dynamodb.CreateTableOutput{},
			stubError:         errors.New("some error"),
			stubDescribeError: &types.ResourceNotFoundException{},
			wantError:         errors.New("some error"),
		},
		{
			name:              "should return nil for error and not create a table when table exists",
			stubOutput:        &dynamodb.CreateTableOutput{},
			stubError:         nil,
			stubDescribeError: nil,
			wantError:         nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockClient = NewMockClient()
			var db DynamoDB
			db.Client = mockClient

			mockClient.On("DescribeTable", mock.Anything, mock.Anything, mock.Anything).
				Return(&dynamodb.DescribeTableOutput{}, tt.stubDescribeError).Once()
			mockClient.On("DescribeTable", mock.Anything, mock.Anything, mock.Anything).
				Return(&dynamodb.DescribeTableOutput{
					Table: &types.TableDescription{
						TableStatus: types.TableStatusActive,
					},
				}, nil).Once() // Second mock for waiter
			mockClient.On("CreateTable", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.stubOutput, tt.stubError)

			err := db.CreateTable("table-name", &dynamodb.CreateTableInput{})

			if tt.stubDescribeError != nil {
				mockClient.AssertNumberOfCalls(t, "CreateTable", 1)
			}

			if !reflect.DeepEqual(err, tt.wantError) {
				t.Errorf("Expected error to be nil, but got %s", err)
			}
		})
	}
}

func TestDynamoDB_ScanItem(t *testing.T) {
	tests := []struct {
		name       string
		scanOutput *dynamodb.ScanOutput
		stubError  error
		wantError  error
	}{
		{
			name:       "should return nil for error",
			scanOutput: &dynamodb.ScanOutput{},
			stubError:  nil,
			wantError:  nil,
		},
		{
			name:       "should return error",
			scanOutput: &dynamodb.ScanOutput{},
			stubError:  errors.New("some error"),
			wantError:  errors.New("some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockClient = NewMockClient()
			var db DynamoDB
			db.Client = mockClient

			mockClient.On("Scan", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.scanOutput, tt.stubError)

			err := db.Scan(&dynamodb.ScanInput{})

			mockClient.AssertNumberOfCalls(t, "Scan", 1)

			if !reflect.DeepEqual(err, tt.wantError) {
				t.Errorf("Expected error to be nil, but got %s", err)
			}
		})
	}
}
