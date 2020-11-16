package mock

import "github.com/florestario/core/entity"

type aquaticPlantPersistenceMock struct {
	saveResult error
	getOneReturn *entity.AquaticPlant
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
