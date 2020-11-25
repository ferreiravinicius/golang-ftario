package controller

import (
	"encoding/json"
	entity2 "github.com/florestario/core/entity"
	"github.com/florestario/delivery/api"
	"github.com/florestario/delivery/interactor"
	"net/http"
)

type GenusInput struct {
	Name string `json:"name"`
}

var RegisterGenus api.Handler = func (w http.ResponseWriter, r *http.Request) error {
	var genusInput GenusInput
	if err := Decode(r, &genusInput); err != nil {
		return err
	}

	service := interactor.GenusService()
	entity := &entity2.Genus{ Name: genusInput.Name }
	if output, err := service.CreateGenus(entity); err != nil {
		return err
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&output)
	}
	return nil
}
