package entity_test

import (
	"testing"

	"github.com/florestario/core/entity"
)

func TestAquaticPlant(t *testing.T) {

	genus := &entity.Genus{
		Name: "GenusName",
	}

	specie := &entity.Specie{
		Name: "SpecieName",
		Genus: genus,
	}

	_ = &entity.AquaticPlant{
		ID: 12,
		Variety: "variety",
		Specie: specie,
		BioCycle: entity.Perennial,
	}
}