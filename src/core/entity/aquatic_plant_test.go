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
		Code: "code",
		Variety: "variety",
		Specie: specie,
		BioCycle: entity.Perennial,
	}
}

func Test_UniqueCode(t *testing.T) {
	input := &entity.AquaticPlant{
		Variety: "Variety",
		Specie: &entity.Specie{
			Name: "Specie",
			Genus: &entity.Genus{ Name: "Genus" },
		},
	}

	code, _ := input.UniqueCode()

	if len(code) == 0 {
		t.Errorf("expected code to be generated")
	}

	expectedCode := "GENUSPECVARI"
	if code != expectedCode {
		t.Errorf("expected code to be '%s' and got '%s'", expectedCode, code)
	}

	wrongInput := &entity.AquaticPlant{}
	_, err := wrongInput.UniqueCode()
	if err == nil {
		t.Errorf("expected error when sending empty aquatic plant")
	}
}