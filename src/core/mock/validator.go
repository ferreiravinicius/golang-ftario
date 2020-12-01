package mock

import "github.com/florestario/core/entity"

type ValidatorMock struct {}

func NewValidatorMock() *ValidatorMock {
	return &ValidatorMock{}
}

func (v ValidatorMock) ValidateGenus(_ *entity.Genus) error {
	return nil
}

