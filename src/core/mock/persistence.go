package mock

import (
	"fmt"
	"github.com/florestario/core/entity"
)

type GenusPersistenceMock struct {
	doError bool
}

func NewGenusPersistenceMock(doError bool) *GenusPersistenceMock {
	return &GenusPersistenceMock{doError: doError}
}

func (mock GenusPersistenceMock) SaveGenus(genus *entity.Genus) (*entity.Genus, error) {
	if mock.doError {
		return nil, fmt.Errorf("mock error")
	}
	genus.ID = 555
	return genus, nil
}

type aquaticPlantPersistenceMock struct {
	saveResult error
	getOneReturn *entity.AquaticPlant
	giveError bool
}

func (mock aquaticPlantPersistenceMock) GetOne(specie *entity.Specie, variety string) (*entity.AquaticPlant, error) {
	return mock.getOneReturn, nil
}

func (mock aquaticPlantPersistenceMock) Save(plant *entity.AquaticPlant) error {
	if mock.saveResult == nil {
		plant.ID = 123
	}
	return mock.saveResult
}

// NewAquaticPlantPersistenceMock just a mock
func NewAquaticPlantPersistenceMock(saveResult error, getOneReturn *entity.AquaticPlant) *aquaticPlantPersistenceMock {
	return &aquaticPlantPersistenceMock{
		saveResult: saveResult,
		getOneReturn: getOneReturn,
	}
}
