package showroom

import (
	"github.com/florestario/core/entity"
	"github.com/florestario/core/gateway"
	"strings"
)

type FilterSpecieByNameInteractor struct {
	persistence gateway.SpecieReader
}

func NewFilterSpecieByNameInteractor(persistence gateway.SpecieReader) *FilterSpecieByNameInteractor {
	return &FilterSpecieByNameInteractor{persistence: persistence}
}

func (interactor FilterSpecieByNameInteractor) Execute(name string, genusId ...int64) ([]entity.Specie, error) {

	nameSanitized := strings.TrimSpace(name)
	if len(nameSanitized) == 0 {
		return make([]entity.Specie, 0), nil
	}

	if len(genusId) > 0 && genusId[0] > 0 {
		currentId := genusId[0]
		output, err := interactor.persistence.FilterSpecieByNameAndGenus(nameSanitized, currentId)
		if err != nil {
			return nil, err
		}
		return output, nil
	}

	output, err := interactor.persistence.FilterSpecieByName(nameSanitized)
	if err != nil {
		return nil, err
	}

	return output, nil
}
