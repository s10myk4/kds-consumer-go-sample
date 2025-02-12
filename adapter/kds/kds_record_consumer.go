package kds

import (
	"github.com/aws/aws-lambda-go/events"
	"kds-consumer-go-sample/app/interfaces"
	"kds-consumer-go-sample/app/model"
	"log"
)

type KDSRecordConsumer struct{}

func NewKDSRecordConsumer() interfaces.RecordConsumer[events.KinesisEvent, model.Record] {
	return &KDSRecordConsumer{}
}

func (kds *KDSRecordConsumer) Consume(in *events.KinesisEvent) (*[]model.Record, error) {
	var records []model.Record

	for _, r := range in.Records {
		kinesisRecord := r.Kinesis
		data, err := model.ConvertClientReloadDirectionCommand(kinesisRecord)
		if err != nil {
			return nil, err
		}
		record := model.Record{PartitionKey: kinesisRecord.PartitionKey, SequenceNumber: kinesisRecord.SequenceNumber, Data: *data}
		log.Printf("Record data: %+v\n", record.Data)
		records = append(records, record)
	}

	return &records, nil
}
