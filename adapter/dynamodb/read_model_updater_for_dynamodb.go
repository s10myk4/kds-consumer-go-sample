package _dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"kds-consumer-go-sample/app/interfaces"
	"kds-consumer-go-sample/app/model"
	"strconv"
)

type ReadModelUpdaterForDynamodb struct {
	dynamodbClient *dynamodb.Client
	tableName      string
}

func NewReadModelUpdaterForDynamodb(client *dynamodb.Client, tableName string) interfaces.ReadModelUpdater {
	return &ReadModelUpdaterForDynamodb{client, tableName}
}

func (r *ReadModelUpdaterForDynamodb) Update(in *[]model.ReadModel) error {
	var transactItems []types.TransactWriteItem
	for _, rm := range *in {
		item := types.TransactWriteItem{
			Put: &types.Put{
				TableName: aws.String(r.tableName),
				Item: map[string]types.AttributeValue{
					"id":   &types.AttributeValueMemberN{Value: strconv.FormatUint(uint64(rm.AccountId), 10)},
					"name": &types.AttributeValueMemberS{Value: rm.Name},
				},
			},
		}
		transactItems = append(transactItems, item)
	}
	_, err := r.dynamodbClient.TransactWriteItems(context.Background(), &dynamodb.TransactWriteItemsInput{
		TransactItems: transactItems,
	})
	return err
}
