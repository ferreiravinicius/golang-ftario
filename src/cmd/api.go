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

	api := api.NewApi()
	api.Start()
}
