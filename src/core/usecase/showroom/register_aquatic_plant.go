package showroom

import (
	"fmt"
	"github.com/florestario/core/entity"
	"github.com/florestario/core/err"
	"github.com/florestario/core/gateway"
)

// RegisterAquaticPlantUseCase interface for use case
type RegisterAquaticPlantUseCase interface {
	Execute(*entity.AquaticPlant) (string, error)
}

type registerAquaticPlant struct {
	persistence gateway.AquaticPlantPersistence
}

// NewRegisterAquaticPlant regiter aquatic plant implementation
func NewRegisterAquaticPlant(persistence gateway.AquaticPlantPersistence) *registerAquaticPlant {
	return &registerAquaticPlant{
		persistence: persistence,
	}
}

// Execute the usecase
func (usecase *registerAquaticPlant) Execute(plant *entity.AquaticPlant) (string, error) {

	//TODO: call validator for required fields

	persisted, _ := usecase.persistence.GetOne(plant.Specie, plant.Variety)
	if persisted != nil {
		return "", err.ErrTaxonomyAlreadyExists
	}

	code, err := plant.UniqueCode()
	if err != nil {
		return "", err
	}
	plant.Code = code

	if err = usecase.persistence.Save(plant); err != nil {
		return "", fmt.Errorf("failed to save:  %w", err)
	}

	return code, nil
}