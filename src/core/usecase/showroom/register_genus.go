package showroom

import (
	"github.com/florestario/core/entity"
	"github.com/florestario/core/gateway"
	"github.com/pkg/errors"
)

type RegisterGenusInteractor struct {
	persistence gateway.GenusPersistence
	validator   gateway.ValidatorGateway
}

func NewRegisterGenusInteractor(
	persistence gateway.GenusPersistence,
	validator gateway.ValidatorGateway) *RegisterGenusInteractor {
	return &RegisterGenusInteractor{persistence: persistence, validator: validator}
}

func (interactor RegisterGenusInteractor) Execute(genus *entity.Genus) (*entity.Genus, error) {

	if err := interactor.validator.ValidateGenus(*genus); err != nil {
		return nil, err
	}

	if persisted, _ := interactor.persistence.GetGenusByName(genus.Name); persisted != nil {
		err := errors.New("Genus already exists") //TODO: use custom error to help intl
		return nil, err
	}

	genusOutput, err := interactor.persistence.SaveGenus(genus)
	if err != nil {
		wrapped := errors.Wrap(err, "Error at register genus interactor") //TODO: use custom err to help intl
		return nil, wrapped
	}

	return genusOutput, nil
}
