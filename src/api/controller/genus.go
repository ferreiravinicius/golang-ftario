package controller

import (
	"encoding/json"
	"github.com/florestario/api/engine"
	"github.com/florestario/core/entity"
	"github.com/florestario/core/service"
	"net/http"
)

type genusController struct {
	service service.GenusService
}

func NewGenusController(service service.GenusService) *genusController {
	return &genusController{service: service}
}

type GenusInput struct {
	Name string `json:"name"`
}

func (ctr *genusController) HandleCreateGenus(w http.ResponseWriter, r *http.Request) error {
	var genusInput GenusInput
	if err := engine.Decode(r, &genusInput); err != nil {
		return err
	}

	output, err := ctr.service.CreateGenus(&entity.Genus{Name: genusInput.Name})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&output)

	return nil
}

//var RegisterGenus api.Handler = func (w http.ResponseWriter, r *http.Request) error {
//	var genusInput GenusInput
//	if err := Decode(r, &genusInput); err != nil {
//		return err
//	}
//
//	service := interactor.GenusService()
//	entity := &entity2.Genus{ Name: genusInput.Name }
//	if output, err := service.CreateGenus(entity); err != nil {
//		return err
//	} else {
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusCreated)
//		json.NewEncoder(w).Encode(&output)
//	}
//	return nil
//}
