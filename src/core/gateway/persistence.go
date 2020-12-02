package gateway

import "github.com/florestario/core/entity"

// PlantPersistence aquatic plant persistence gateway
type PlantPersistence interface {
	SavePlant(*entity.AquaticPlant) error
	GetPlant(specie *entity.Specie, variety string) (*entity.AquaticPlant, error)
}

// GenusPersistence Genus persistence gateway
type GenusPersistence interface {
	SaveGenus(genus *entity.Genus) (*entity.Genus, error)
	GetGenusByName(name string) (*entity.Genus, error)
}

type ShowroomPersistence interface {
	GenusPersistence
	PlantPersistence
}
