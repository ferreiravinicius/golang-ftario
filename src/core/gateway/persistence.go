package gateway

import "github.com/florestario/core/entity"

// PlantPersistence Aquatic plant persistence gateway
type PlantPersistence interface {
	SavePlant(*entity.AquaticPlant) error
}


// GenusReader Genus persistence sub interface (writer)
type GenusWriter interface {
	SaveGenus(genus *entity.Genus) (*entity.Genus, error)
}

// GenusReader Genus persistence sub interface (reader)
type GenusReader interface {
	GetGenusByName(name string) (*entity.Genus, error)
	FilterGenusByName(name string) ([]entity.Genus, error)
}

type SpeciePersistence interface {
	SpecieReader
}

// GenusPersistence Genus persistence gateway
type GenusPersistence interface {
	GenusReader
	GenusWriter
}

// ShowroomPersistence All persistence gateways grouped
type ShowroomPersistence interface {
	GenusPersistence
	PlantPersistence
	SpecieReader
}
