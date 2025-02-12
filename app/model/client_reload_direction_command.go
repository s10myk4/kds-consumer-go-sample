package model

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/moznion/go-optional"
)

type ClientReloadDirectionCommand struct {
	SourceGithubRepositoryName string                  `json:"source_github_repository_name"`
	AccountId                  uint                    `json:"account_id"`
	TargetId                   optional.Option[string] `json:"target_id"`
	TargetInId                 optional.Option[string] `json:"target_in_id"`
	Type                       string                  `json:"type"`
	Act                        optional.Option[string] `json:"act"`
}

func ConvertClientReloadDirectionCommand(record events.KinesisRecord) (*ClientReloadDirectionCommand, error) {
	var model ClientReloadDirectionCommand
	err := json.Unmarshal(record.Data, &model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}
