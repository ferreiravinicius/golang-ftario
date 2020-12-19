package showroom

import (
	"github.com/florestario/core/entity"
	"github.com/florestario/core/gateway"
	"strings"
)

type FilterGenusByNameInteractor struct {
	persistence gateway.GenusReader
}

func NewFilterGenusByNameInteractor(persistence gateway.GenusReader) *FilterGenusByNameInteractor {
	return &FilterGenusByNameInteractor{persistence: persistence}
}

func (interactor FilterGenusByNameInteractor) Execute(name string) ([]entity.Genus, error) {

	nameSanitized := strings.TrimSpace(name)
	if len(nameSanitized) == 0 {
		return make([]entity.Genus, 0), nil
	}

	output, err := interactor.persistence.FilterGenusByName(nameSanitized)
	if err != nil {
		return nil, err
	}

	return output, nil
}
