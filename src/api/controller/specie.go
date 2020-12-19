package controller

import (
	"encoding/json"
	"github.com/florestario/api/engine"
	"github.com/florestario/core/usecase/showroom"
	"net/http"
	"strconv"
	"strings"
)

type SpecieOutput struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Genus *GenusOutput
}

func HandlerFilterSpecieByName(interactor *showroom.FilterSpecieByNameInteractor) engine.AppHttpHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		name := r.FormValue("name")
		genusId := r.FormValue("genusId")

		currentId, _ := strconv.ParseInt(strings.TrimSpace(genusId), 10, 64)

		filterOutput, err := interactor.Execute(name, currentId)
		if err != nil {
			return err
		}

		output := make([]SpecieOutput, 0, 10)
		for _, item := range filterOutput {
			genus := &GenusOutput{ID: item.Genus.ID, Name: item.Genus.Name}
			specie := SpecieOutput{ID: item.ID, Name: item.Name, Genus: genus}
			output = append(output, specie)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&output)
		return nil
	}
}
