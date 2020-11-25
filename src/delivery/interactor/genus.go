package interactor

import (
	"github.com/florestario/core/service"
	"github.com/florestario/persistence"
)

func GenusPersistence() *service.GenusPersistence {
	var persistence service.GenusPersistence = persistence.NewAquaticPostgres()
	return &persistence
}

func GenusService() service.GenusService {
	persistence := GenusPersistence()
	var service service.GenusService = service.NewGenusService(*persistence)
	return service
}
