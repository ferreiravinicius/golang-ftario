package validator

import (
	"github.com/florestario/core/entity"
	"github.com/florestario/core/err"
)

var (
	ErrGenusValidation = err.New("ErrGenusVal", "Genus is required and can't be empty")
	ErrSpecieNameValidation = err.New("ErrSpecieNameVal", "Specie name is required")
)

func ValidateGenus(genus *entity.Genus) error {

	if genus == nil {
		errMessage := "Genus is required"
		return err.New("ErrGenusNil", errMessage)
	}

	if genus.Name == "" {
		return ErrGenusValidation
	}
	return nil
}

func ValidateSpecie(specie *entity.Specie) error {

	if specie == nil {
		errMessage := "Specie is required"
		return err.New("ErrSpecNil", errMessage)
	}

	if specie.Name == "" {
		errMessage, errArg := "Specie '%s' is required", "name"
		return err.New("ErrSpecNameVal", errMessage, errArg)
	}

	if err := ValidateGenus(specie.Genus); err != nil {
		return err
	}

	return nil
}


