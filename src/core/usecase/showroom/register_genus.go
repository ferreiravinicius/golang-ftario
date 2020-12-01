package showroom

import (
	"github.com/florestario/core/entity"
	"github.com/florestario/core/gateway"
)

type registerGenusInteractor struct {
	persistence gateway.GenusPersistence
	validator   gateway.ValidatorGateway
}

func NewRegisterGenusInteractor(
	persistence gateway.GenusPersistence,
	validator gateway.ValidatorGateway) *registerGenusInteractor {
	return &registerGenusInteractor{persistence: persistence, validator: validator}
}

func (interactor registerGenusInteractor) Execute(genus *entity.Genus) (*entity.Genus, error) {

	if err := interactor.validator.ValidateGenus(genus); err != nil {
		return nil, err
	}

	genusOutput, err := interactor.persistence.SaveGenus(genus)
	if err != nil {
		//TODO: validate if message provider is necessary for this kind of error
		return nil, err
	}

	return genusOutput, nil
}
