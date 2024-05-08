package database

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) CreateTable(ctx context.Context, input *dynamodb.CreateTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error) {
	args := m.Called(ctx, input, optFns)
	return args.Get(0).(*dynamodb.CreateTableOutput), args.Error(1)
}

func (m *MockClient) DescribeTable(ctx context.Context, params *dynamodb.DescribeTableInput, f ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error) {
	args := m.Called(ctx, params, f)
	return args.Get(0).(*dynamodb.DescribeTableOutput), args.Error(1)
}

func (m *MockClient) DeleteTable(ctx context.Context, input *dynamodb.DeleteTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteTableOutput, error) {
	args := m.Called(ctx, input, optFns)
	return args.Get(0).(*dynamodb.DeleteTableOutput), args.Error(1)
}

func (m *MockClient) PutItem(ctx context.Context, input *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	args := m.Called(ctx, input, optFns)
	return args.Get(0).(*dynamodb.PutItemOutput), args.Error(1)
}

func NewMockClient() *MockClient {
	return &MockClient{}
}
