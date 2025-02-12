package model

type ReadModel struct {
	AccountId uint   `json:"id"`
	Name      string `json:"name"`
}

func ConvertReadModel(command ClientReloadDirectionCommand) *ReadModel {
	return &ReadModel{
		command.AccountId,
		command.Type,
	}
}
