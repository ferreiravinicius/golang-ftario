package validator

import (
	"github.com/florestario/core/entity"
	"github.com/florestario/core/errx"
)

var (
	ErrGenusValidation = errx.New("ErrGenusVal", "Genus is required and can't be empty")
	ErrSpecieNameValidation = errx.New("ErrSpecieNameVal", "Specie name is required")
)

func ValidateGenus(genus *entity.Genus) error {

	if genus == nil {
		errMessage := "Genus is required"
		return errx.New("ErrGenusNil", errMessage)
	}

	if genus.Name == "" {
		return ErrGenusValidation
	}
	return nil
}

func ValidateSpecie(specie *entity.Specie) error {

	if specie == nil {
		errMessage := "Specie is required"
		return errx.New("ErrSpecNil", errMessage)
	}

	if specie.Name == "" {
		errMessage, errArg := "Specie '%s' is required", "name"
		return errx.New("ErrSpecNameVal", errMessage, errArg)
	}

	if err := ValidateGenus(specie.Genus); err != nil {
		return err
	}

	return nil
}


