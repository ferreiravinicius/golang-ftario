package validator

import (
	"github.com/florestario/core/entity"
	"github.com/florestario/core/gateway"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"strings"
)

type ozzoValidator struct {
}

func NewOzzoValidator() gateway.ValidatorGateway {
	return &ozzoValidator{}
}

// Validate Basic validation of genus fields
func (o ozzoValidator) ValidateGenus(genus entity.Genus) error {
	sanitize(&genus)
	err := validation.ValidateStruct(&genus,
		validation.Field(&genus.Name, validation.Required, validation.Length(2, 50)),
	)
	return err
}

// Sanitizes fields that can be spaces and etc
func sanitize(genus *entity.Genus) {
	genus.Name = strings.TrimSpace(genus.Name)
}
