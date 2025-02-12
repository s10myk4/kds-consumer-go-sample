package app

import (
	"github.com/aws/aws-lambda-go/events"
	"kds-consumer-go-sample/app/interfaces"
	"kds-consumer-go-sample/app/model"
	"log"
)

type App interface {
	Run(in events.KinesisEvent) error
}

type app struct {
	recordConsumer   interfaces.RecordConsumer[events.KinesisEvent, model.Record]
	readModelUpdater interfaces.ReadModelUpdater
}

func NewApp(recordConsumer interfaces.RecordConsumer[events.KinesisEvent, model.Record], readModelUpdater interfaces.ReadModelUpdater) App {
	return &app{recordConsumer, readModelUpdater}
}

func (app *app) Run(in events.KinesisEvent) error {
	records, err := app.recordConsumer.Consume(&in)
	if err != nil {
		return err
	}
	var readModels []model.ReadModel
	for _, r := range *records {
		rm := model.ConvertReadModel(r.Data)
		readModels = append(readModels, *rm)
	}
	err = app.readModelUpdater.Update(&readModels)
	if err != nil {
		log.Printf("Failed consume: %v\n", err)
		return err
	}
	log.Println("Successful consume")
	return nil
}
