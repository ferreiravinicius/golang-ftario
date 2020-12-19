package gateway

import "github.com/florestario/core/entity"

// SpecieReader Specie persistence sub interface (reader)
type SpecieReader interface {
	FilterSpecieByName(name string) ([]entity.Specie, error)
	FilterSpecieByNameAndGenus(name string, genusId int64) ([]entity.Specie, error)
}
