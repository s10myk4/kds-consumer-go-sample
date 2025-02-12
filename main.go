package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"kds-consumer-go-sample/adapter/kds"
	"kds-consumer-go-sample/app"
	"kds-consumer-go-sample/app/interfaces"
	"log"
)

var dynamodbClient *dynamodb.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	dynamodbClient = dynamodb.NewFromConfig(cfg)
}

// https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/golang-handler.html#golang-handler-naming
func handleRequest(ctx context.Context, event events.KinesisEvent) error {
	//tableName := os.Getenv("DYNAMODB_TABLE_NAME")
	//if tableName == "" {
	//	log.Fatal("missing environment variable DYNAMODB_TABLE_NAME")
	//}

	app := app.NewApp(
		kds.NewKDSRecordConsumer(),
		interfaces.NewMockReadModelUpdater(),
		//_dynamodb.NewReadModelUpdaterForDynamodb(dynamodb.NewFromConfig(cfg), tableName),
	)
	err := app.Run(event)
	return err
}

func main() {
	lambda.Start(handleRequest)
}
