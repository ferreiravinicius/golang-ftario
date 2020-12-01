package main

import (
	"github.com/florestario/api"
)


//type Customer struct {
//	Name    string
//	Gender  string
//	Email   string
//}

//func testingValidator() {
//	c := Customer{
//		Name:  "Qi",
//		Email: "q",
//	}
//
//	err := validation.Errors{
//		"name": validation.Validate(c.Name, validation.Required, validation.Length(5, 20).Error("test {{.min}} and {1}")),
//		"email": validation.Validate(c.Name, validation.Required, is.Email),
//	}.Filter()
//	fmt.Println(err)
//}

func main() {

	//testingValidator()

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
