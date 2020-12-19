package controller

import (
	"encoding/json"
	"github.com/florestario/api/engine"
	"github.com/florestario/core/entity"
	"github.com/florestario/core/usecase/showroom"
	"net/http"
)

type GenusInput struct {
	Name string `json:"name"`
}

type GenusOutput struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func HandlerRegisterGenus(interactor *showroom.RegisterGenusInteractor) engine.AppHttpHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var genusInput GenusInput
		if err := engine.Decode(r, &genusInput); err != nil {
			return err
		}

		genusPersisted, err := interactor.Execute(&entity.Genus{Name: genusInput.Name})
		if err != nil {
			return err
		}

		output := &GenusOutput{
			ID:   genusPersisted.ID,
			Name: genusPersisted.Name,
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&output)

		return nil
	}
}

func HandlerFilterGenusByName(interactor *showroom.FilterGenusByNameInteractor) engine.AppHttpHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		name := r.FormValue("name")

		filterOutput, err := interactor.Execute(name)
		if err != nil {
			return err
		}

		output := make([]GenusOutput, 0, 10)
		for _, item := range filterOutput {
			genus := GenusOutput{ID: item.ID, Name: item.Name}
			output = append(output, genus)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&output)
		return nil
	}
}
