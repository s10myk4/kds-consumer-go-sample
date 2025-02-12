package model

type Record struct {
	PartitionKey   string
	SequenceNumber string
	Data           ClientReloadDirectionCommand
}
