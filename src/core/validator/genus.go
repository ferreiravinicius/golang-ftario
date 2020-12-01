package validator

import (
	"github.com/florestario/core/entity"
	"github.com/florestario/core/err"
	"github.com/florestario/core/gateway"
	"strings"
)

type GenusValidator interface {
	Must(genus *entity.Genus) error
}

type IntlGenusValidator struct {
	message gateway.MessageProvider
}

func (validator *IntlGenusValidator) Must(genus *entity.Genus) error {

	if genus == nil {
		return err.Basic("genus.name.required", "Genus requires a valid name")
	}

	name := strings.TrimSpace(genus.Name)
	if len(name) < 2 {
		return err.Basic("genus.name.required", "Genus requires a valid name")
	}

	return nil
}

type CustomVal interface {
	Required(name string)
}
