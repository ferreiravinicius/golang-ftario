package gateway

import "github.com/florestario/core/entity"

type GenusPersistence interface {
	SaveGenus(genus *entity.Genus) (*entity.Genus, error)
}
