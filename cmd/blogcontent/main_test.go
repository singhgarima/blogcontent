package main

import (
	"reflect"
	"testing"
)

func Test_loadDynamoDBConfig(t *testing.T) {
	tests := []struct {
		name       string
		awsRegion  string
		wantRegion string
		wantErr    error
	}{
		{
			name:       "Should set config with local db for AWS Region as localhost",
			awsRegion:  "localhost",
			wantRegion: "localhost",
			wantErr:    nil,
		}, {
			name:       "Should set config with AWS db for AWS Region as ap-southeast-4",
			awsRegion:  "ap-southeast-4",
			wantRegion: "ap-southeast-4",
			wantErr:    nil,
		}, {
			name:       "Should set config with AWS db for AWS Region as ap-southeast-2 when not set",
			awsRegion:  "",
			wantRegion: "ap-southeast-2",
			wantErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("AWS_REGION", tt.awsRegion)
			got, err := loadDynamoDBConfig()

			if err != tt.wantErr {
				t.Errorf("loadDynamoDBConfig() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got.Region, tt.wantRegion) {
				t.Errorf("loadDynamoDBConfig() got region = %v, want region %v", got.Region, tt.wantRegion)
			}

		})
	}
}
