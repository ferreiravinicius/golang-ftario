package mock

import (
	"fmt"
	"github.com/florestario/core/entity"
)

type GenusPersistenceMock struct {
	MockError          bool
	MockSaveGenusId    int
	MockGetGenusByName *entity.Genus
}

func (mock GenusPersistenceMock) GetGenusByName(name string) (*entity.Genus, error) {
	return mock.MockGetGenusByName, nil
}

func (mock GenusPersistenceMock) SaveGenus(genus *entity.Genus) (*entity.Genus, error) {
	if mock.MockError {
		return nil, fmt.Errorf("mock error")
	}
	genus.ID = mock.MockSaveGenusId
	return genus, nil
}

type aquaticPlantPersistenceMock struct {
	saveResult   error
	getOneReturn *entity.AquaticPlant
	giveError    bool
}

func (mock aquaticPlantPersistenceMock) GetPlant(specie *entity.Specie, variety string) (*entity.AquaticPlant, error) {
	return mock.getOneReturn, nil
}

func (mock aquaticPlantPersistenceMock) SavePlant(plant *entity.AquaticPlant) error {
	if mock.saveResult == nil {
		plant.ID = 123
	}
	return mock.saveResult
}

// NewAquaticPlantPersistenceMock just a mock
func NewAquaticPlantPersistenceMock(saveResult error, getOneReturn *entity.AquaticPlant) *aquaticPlantPersistenceMock {
	return &aquaticPlantPersistenceMock{
		saveResult:   saveResult,
		getOneReturn: getOneReturn,
	}
}
