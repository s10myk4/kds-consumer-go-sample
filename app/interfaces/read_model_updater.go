package interfaces

import (
	"kds-consumer-go-sample/app/model"
)

type ReadModelUpdater interface {
	Update(in *[]model.ReadModel) error
}

type MockReadModelUpdater struct{}

func NewMockReadModelUpdater() ReadModelUpdater {
	return &MockReadModelUpdater{}
}

func (r *MockReadModelUpdater) Update(in *[]model.ReadModel) error {
	return nil
}
