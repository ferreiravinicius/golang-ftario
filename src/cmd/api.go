package main

import (
	"github.com/florestario/api"
)

func main() {

	//pg := persistence.Instance()
	//r, err := pg.FilterSpecieByNameAndGenus("esx", 14)
	//
	//fmt.Println(r)
	//fmt.Println(err)

	api := api.NewApi()
	api.Start()
}
