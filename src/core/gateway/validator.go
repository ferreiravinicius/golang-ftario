package gateway

import "github.com/florestario/core/entity"

type ValidatorGateway interface {
	ValidateGenus(genus *entity.Genus) error
}
