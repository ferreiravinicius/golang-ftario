package service

import (
	"github.com/florestario/core/entity"
	"github.com/florestario/core/errx"
	"github.com/florestario/core/validator"
)

type GenusService interface {
	CreateGenus(genus *entity.Genus) (*entity.Genus, error)
}

type GenusPersistence interface {
	SaveGenus(genus *entity.Genus) (*entity.Genus, error)
}

type genusService struct {
	persistence GenusPersistence
}

func NewGenusService(persistence GenusPersistence) *genusService {
	return &genusService{persistence: persistence}
}

func (service *genusService) CreateGenus(genus *entity.Genus) (*entity.Genus, error) {

	if err := validator.ValidateGenus(genus); err != nil {
		return nil, err
	}

	genusOutput, err := service.persistence.SaveGenus(genus)
	if err != nil {
		errorMessage := "Unexpected error: %w"
		errx.New("ErrUxp", errorMessage, err)
	}

	return genusOutput, nil
}



