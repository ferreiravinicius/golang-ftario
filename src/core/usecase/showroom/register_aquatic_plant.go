package showroom

import (
	"fmt"
	"github.com/florestario/core/entity"
	"github.com/florestario/core/err"
	"github.com/florestario/core/gateway"
)

// RegisterAquaticPlantUseCase interface for use case
type RegisterAquaticPlantUseCase interface {
	Execute(*entity.AquaticPlant) (*entity.AquaticPlant, error)
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
func (usecase *registerAquaticPlant) Execute(plant *entity.AquaticPlant) (*entity.AquaticPlant, error) {

	//TODO: call validator for required fields

	persisted, _ := usecase.persistence.GetOne(plant.Specie, plant.Variety)
	if persisted != nil {
		return nil, err.ErrTaxonomyAlreadyExists
	}

	if err := usecase.persistence.Save(plant); err != nil {
		return nil, fmt.Errorf("failed to save:  %w", err)
	}

	return plant, nil
}