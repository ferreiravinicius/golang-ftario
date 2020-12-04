package validator_test

import (
	"github.com/florestario/adapter/validator"
	"github.com/florestario/core/entity"
	"github.com/florestario/core/gateway"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"testing"
)

func TestOzzoValidator_ValidateGenus(t *testing.T) {

	var validatorGateway gateway.ValidatorGateway
	validatorGateway = validator.NewOzzoValidator()

	inputs := map[string]entity.Genus{
		"genusEmptyName":      {Name: ""},
		"genusOneLetterName":  {Name: "a"},
		"genusOnlySpacesName": {Name: "      "},
	}

	for name, genus := range inputs {
		err := validatorGateway.ValidateGenus(genus)
		if err == nil {
			t.Errorf("Expected error at %s", name)
		}
		if _, ok := err.(validation.Errors); !ok {
			t.Errorf("Exptected error type to be validation error")
		}
	}
}
