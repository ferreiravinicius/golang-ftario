package showroom

import (
	"github.com/florestario/core/entity"
	"github.com/florestario/core/mock"
	"testing"
)

func TestExecute(t *testing.T) {

	t.Run("should return saved genus", func(t *testing.T) {

		expectedId := 999

		persistence := mock.GenusPersistenceMock{
			MockSaveGenusId: expectedId,
		}
		validator := mock.ValidatorMock{}
		interactor := NewRegisterGenusInteractor(persistence, validator)

		output, _ := interactor.Execute(&entity.Genus{
			Name: "test",
		})

		if output.ID != expectedId {
			t.Errorf("expected ID to be fullfilled with persisted data")
		}
	})

	t.Run("should check if exists genus with same name", func(t *testing.T) {

		fakeGenus := &entity.Genus{Name: "fake"}
		persistence := mock.GenusPersistenceMock{ MockGetGenusByName: fakeGenus }
		validator := mock.ValidatorMock{}
		interactor := NewRegisterGenusInteractor(persistence, validator)

		if _, err := interactor.Execute(&entity.Genus{}); err == nil {
			t.Errorf("expected error when duplicated is found")
		}
	})
}
