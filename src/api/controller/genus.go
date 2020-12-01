package controller

import (
	"encoding/json"
	"github.com/florestario/api/engine"
	"github.com/florestario/core/entity"
	"github.com/florestario/core/service"
	"github.com/florestario/core/usecase/showroom"
	"net/http"
)

type GenusController struct {
	interactor showroom.RegisterGenusInteractor
}

func NewGenusController(interactor showroom.RegisterGenusInteractor) *GenusController {
	return &GenusController{interactor: interactor}
}

type GenusInput struct {
	Name string `json:"name"`
}

func (ctr *GenusController) HandleCreateGenus(w http.ResponseWriter, r *http.Request) error {
	var genusInput GenusInput
	if err := engine.Decode(r, &genusInput); err != nil {
		return err
	}

	output, err := ctr.interactor.Execute(&entity.Genus{Name: genusInput.Name})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&output)

	return nil
}
