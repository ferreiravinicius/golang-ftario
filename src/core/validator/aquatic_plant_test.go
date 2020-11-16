package validator_test

import (
	"github.com/florestario/core/entity"
	"github.com/florestario/core/validator"
	"testing"
)

func TestValidateSpecie(t *testing.T) {

	t.Run("specie must contain a valid genus", func(t *testing.T) {

		specie := &entity.Specie{
			Name:  "Whatever",
			Genus: nil,
		}

		if err := validator.ValidateSpecie(specie); err == nil {
			t.Errorf("specie should return error when genus is wrong")
		}
	})

	t.Run("specie must contain a valid name", func(t *testing.T) {

		genus := &entity.Genus{Name: "Genus"}

		wrongSpecie := entity.Specie{Name:  "", Genus: genus}
		if err := validator.ValidateSpecie(&wrongSpecie); err == nil {
			t.Errorf("expected error when name is invalid")
		}

		var nilSpecie *entity.Specie
		if err := validator.ValidateSpecie(nilSpecie); err == nil {
			t.Errorf("expected error when specie is nil")
		}
	})



}

func TestValidateGenus(t *testing.T) {
	var (
		genusEmptyName = entity.Genus{Name: ""}
		genusEmpty = entity.Genus{}
		genusCorrect = entity.Genus{Name: "Whatever name"}
	)

	if err := validator.ValidateGenus(&genusEmptyName); err == nil {
		t.Errorf("should return error when wrong genus is passed")
	}

	if err := validator.ValidateGenus(&genusEmpty); err == nil {
		t.Errorf("should return error when wrong genus is passed")
	}

	if err := validator.ValidateGenus(&genusCorrect); err != nil {
		t.Errorf("valid input should pass")
	}
}