package api

import (
	"github.com/florestario/core/mock"
	"github.com/florestario/core/usecase/showroom"
)

var persistenceMock = mock.NewAquaticPlantPersistenceMock(nil, nil)

var (
	RegisterAquaticPlant = showroom.NewRegisterAquaticPlant(persistenceMock)
)
