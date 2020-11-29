package main

import (
	"github.com/florestario/api"
)


func main() {


	//input := &entity.AquaticPlant{
	//	Variety: "varx2",
	//	Specie: &entity.Specie{
	//		Name: "specx1",
	//		Genus: &entity.Genus{
	//			Name: "genusx1",
	//		},
	//	},
	//}
	//
	//pg := persistence.NewAquaticPostgres()
	//pg.SaveGenus(input)

	//persistence.Other()

	api := api.NewApi()
	api.Start()
}
