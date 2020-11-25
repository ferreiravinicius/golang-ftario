package delivery

import (
	"encoding/json"
	"fmt"
	"github.com/florestario/core/entity"
	"github.com/florestario/core/service"
	"github.com/florestario/core/usecase/showroom"
	"github.com/florestario/persistence"
	"net/http"
)


func Decode(request *http.Request, output interface{}) error {
	if err := json.NewDecoder(request.Body).Decode(&output); err != nil {
		return fmt.Errorf("Something is wrong with input (%w) ", err)
	}
	return nil
}

type GenusOutput struct {
	ID      int    `json:"id"`
	Name  string `json:"specie"`
}

func CreateGenus(w http.ResponseWriter, r *http.Request) {
	var genusInput entity.Genus
	if err := Decode(r, &genusInput); err != nil {
		AnswerError(w, http.StatusBadRequest, "ErrInp", err.Error())
	} else {
		pg := persistence.NewAquaticPostgres()
		genusService := service.NewGenusService(pg)
		output, err := genusService.CreateGenus(&genusInput)
		if err != nil {
			AnswerError(w, http.StatusBadRequest, "ErrNeedToWrap", err.Error())
		}
		AnswerSuccess(w, http.StatusAccepted, output)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	var input AquaticPlantInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		message := fmt.Sprintf("Something wrong with input (%s)", err.Error())
		AnswerError(w, http.StatusBadRequest, "ErrInp", message)
	} else {
		pg := persistence.NewAquaticPostgres()
		usecase := showroom.NewRegisterAquaticPlant(pg)
		plant, err := usecase.Execute(input.ToEntity())
		if err != nil {
			AnswerError(w, http.StatusBadRequest, "ErrNeedToWrap", err.Error())
		}

		output := FromEntity(plant)
		AnswerSuccess(w, http.StatusAccepted, output)
	}
}
