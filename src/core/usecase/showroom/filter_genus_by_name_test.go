package showroom_test

import (
	"github.com/florestario/core/entity"
	"github.com/florestario/core/mock"
	"github.com/florestario/core/usecase/showroom"
	"testing"
)

func TestCurrentUseCase(t *testing.T) {


	t.Run(`should sanitize input to avoid blank name`, func(t *testing.T) {
		fakeOutput := []entity.Genus{{ID: 1010, Name: "Testing"}}
		persistence := &mock.GenusReaderMock{MockFilterGenusByName: fakeOutput}
		interactor := showroom.NewFilterGenusByNameInteractor(persistence)

		output1, err1 := interactor.Execute("")
		if err1 != nil || len(output1) > 0 {
			t.Errorf("expected empty result when given blank name parameter")
		}

		output2, err2 := interactor.Execute("               ")
		if err2 != nil || len(output2) > 0 {
			t.Errorf("expected empty result when given blank name parameter")
		}
	})

}
