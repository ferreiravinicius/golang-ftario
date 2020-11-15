package main

import (
	"github.com/florestario/core/entity"
	"github.com/florestario/persistence"
)


func main() {


	input := &entity.AquaticPlant{
		Code:    "codex1",
		Variety: "varx2",
		Specie: &entity.Specie{
			Name: "specx1",
			Genus: &entity.Genus{
				Name: "genusx1",
			},
		},
	}

	pg := persistence.NewAquaticPostgres()
	pg.Save(input)

//	persistence.Other()

//	api := api.NewApi()
//	api.Start()
}
