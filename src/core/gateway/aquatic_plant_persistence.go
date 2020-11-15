package gateway

import "github.com/florestario/core/entity"

// AquaticPlantPersistence aquatic plant persistence gateway
type AquaticPlantPersistence interface {
	Save(*entity.AquaticPlant) error
	GetOne(specie *entity.Specie, variety string) (*entity.AquaticPlant, error)
}