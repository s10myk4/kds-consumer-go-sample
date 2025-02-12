package model

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/moznion/go-optional"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertClientReloadDirectionCommand(t *testing.T) {

	testCases := []struct {
		name          string
		inputJSON     string
		expectedModel ClientReloadDirectionCommand
		expectError   bool
	}{
		{
			name: "正常系",
			inputJSON: `{
				"source_github_repository_name": "example-repo",
				"account_id": 12345,
				"target_id": "target123",
				"target_in_id": "targetin456",
				"type": "room",
				"act": "chat"
			}`,
			expectedModel: ClientReloadDirectionCommand{
				SourceGithubRepositoryName: "example-repo",
				AccountId:                  12345,
				TargetId:                   optional.Some("target123"),
				TargetInId:                 optional.Some("targetin456"),
				Type:                       "room",
				Act:                        optional.Some("chat"),
			},
			expectError: false,
		},
		{
			name: "optionなフィールドに値が指定されてない場合",
			inputJSON: `{
				"source_github_repository_name": "example-repo",
				"account_id": 12345,
				"type": "room"
			}`,
			expectedModel: ClientReloadDirectionCommand{
				SourceGithubRepositoryName: "example-repo",
				AccountId:                  12345,
				TargetId:                   optional.None[string](),
				TargetInId:                 optional.None[string](),
				Type:                       "room",
				Act:                        optional.None[string](),
			},
			expectError: false,
		},
		{
			name: "不正なJSONデータの場合",
			inputJSON: `{
				"source_github_repository_name": "example-repo",
				"account_id": "invalid-id",
				"type": "room"
			}`,
			expectedModel: ClientReloadDirectionCommand{},
			expectError:   true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			record := events.KinesisRecord{
				Data: []byte(tt.inputJSON),
			}

			result, err := ConvertClientReloadDirectionCommand(record)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedModel, *result)
			}
		})
	}
}
