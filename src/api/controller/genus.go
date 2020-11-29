package controller

import (
	"fmt"
	"github.com/florestario/core/service"
	"net/http"
)

type GenusInput struct {
	Name string `json:"name"`
}

func CreateGenus(service *service.GenusService) Handler {
	return func (w http.ResponseWriter, r *http.Request) error {
		return fmt.Errorf("Testing returned type ")
	}
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
