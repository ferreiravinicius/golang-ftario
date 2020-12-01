package showroom

import (
	"github.com/florestario/core/entity"
	"github.com/florestario/core/mock"
	"testing"
)

func TestExecute(t *testing.T) {

	t.Run("should return saved genus", func(t *testing.T) {

		persistence := mock.NewGenusPersistenceMock(false)
		validator := mock.NewValidatorMock()
		interactor := NewRegisterGenusInteractor(persistence, validator)

		output, _ := interactor.Execute(&entity.Genus{
			Name: "test",
		})

		if output.ID == 0 {
			t.Errorf("expected ID to be fullfilled with persisted data")
		}
	})

}
