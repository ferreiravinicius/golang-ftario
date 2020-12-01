package showroom_test

import (
	"errors"
	"testing"

	"github.com/florestario/core/mock"

	"github.com/florestario/core/entity"
	"github.com/florestario/core/usecase/showroom"
)

func getValidInput() *entity.AquaticPlant {
	return &entity.AquaticPlant{
		Variety: "Variety",
		Specie: &entity.Specie{
			Name: "Specie",
			Genus: &entity.Genus{ Name: "Genus" },
		},
	}
}

func TestPersistenceGateway(t *testing.T) {

	t.Run("should return err when save fails", func(t *testing.T) {

		var errorMock = errors.New("mock err")
		persistenceMock := mock.NewAquaticPlantPersistenceMock(errorMock, nil)

		ucase := showroom.NewRegisterAquaticPlant(persistenceMock)

		_, err := ucase.Execute(getValidInput())
		if !errors.Is(err, errorMock) {
			t.Errorf("expected err when save fails")
		}
	})

	t.Run("should return err when duplicated variety and specie", func(t *testing.T) {
		var plantMock = &entity.AquaticPlant{}
		persistenceMock := mock.NewAquaticPlantPersistenceMock(nil, plantMock)

		ucase := showroom.NewRegisterAquaticPlant(persistenceMock)
		_, err := ucase.Execute(getValidInput())
		if err == nil {
			t.Errorf("expected err when found duplicated aquatic plant")
		}
	})
}

func TestShouldCallCodeGeneratorAndReturnCodeWhenSaved(t *testing.T) {
	persistenceMock := mock.NewAquaticPlantPersistenceMock(nil, nil)
	ucase := showroom.NewRegisterAquaticPlant(persistenceMock)
	plant, _ := ucase.Execute(getValidInput())
	if plant.ID <= 0 {
		t.Errorf("should return generated id when finishes saving")
	}
}