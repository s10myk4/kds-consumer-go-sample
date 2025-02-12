package kds

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/cstanze/stripmargin"
	"github.com/moznion/go-optional"
	"github.com/stretchr/testify/assert"
	"kds-consumer-go-sample/app/model"
	"testing"
)

func TestConsume(t *testing.T) {
	consumer := NewKDSRecordConsumer()

	testCases := []struct {
		name          string
		input         events.KinesisEvent
		expected      []model.Record
		expectedError bool
	}{
		{
			name: "正常系",
			input: events.KinesisEvent{
				Records: []events.KinesisEventRecord{
					{
						Kinesis: events.KinesisRecord{
							PartitionKey:   "partitionKey1",
							SequenceNumber: "12345",
							Data: []byte(stripmargin.StripMargin(`|
                            |{
							|   "source_github_repository_name": "test",
        					|   "account_id": 1,
        					|   "target_id": "2",
        					|   "target_in_id": "3",
        					|   "type": "room",
        					|   "act": "read"
        					|}`)),
						},
					},
				},
			},
			expected: []model.Record{
				{
					PartitionKey:   "partitionKey1",
					SequenceNumber: "12345",
					Data: model.ClientReloadDirectionCommand{
						SourceGithubRepositoryName: "test",
						AccountId:                  1,
						TargetId:                   optional.Some("2"),
						TargetInId:                 optional.Some("3"),
						Type:                       "room",
						Act:                        optional.Some("read"),
					},
				},
			},
			expectedError: false,
		},
		{
			name: "不正なRecord Dataの場合",
			input: events.KinesisEvent{
				Records: []events.KinesisEventRecord{
					{
						Kinesis: events.KinesisRecord{
							PartitionKey:   "partitionKey2",
							SequenceNumber: "67890",
							Data:           []byte(`invalid data`),
						},
					},
				},
			},
			expected:      nil,
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := consumer.Consume(&tc.input)
			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, *result)
			}
		})
	}
}
