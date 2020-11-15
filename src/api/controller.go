package api

import (
	"encoding/json"
	"fmt"
	"github.com/florestario/core/entity"
	"net/http"
)

type AquaticPlantInput struct {
	Code    string `json:"code"`
	Variety string `json:"variety"`
	Specie  string `json:"specie"`
}

func (input *AquaticPlantInput) ToEntity() *entity.AquaticPlant {
	return &entity.AquaticPlant{
		Code:    input.Code,
		Specie:  &entity.Specie{Name: input.Specie},
		Variety: input.Variety,
	}
}

type ErrorOutput struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func AnswerError(w http.ResponseWriter, statusCode int, code, message string) {

	w.Header().Set("Content-Type", "application/json")

	output := ErrorOutput{
		Code:    code,
		Message: message,
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&output)
}

func AnswerSuccess(w http.ResponseWriter, statusCode int, data ...interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if len(data) == 1 {
		json.NewEncoder(w).Encode(&data)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	var input AquaticPlantInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		message := fmt.Sprintf("Something wrong with input (%s)", err.Error())
		AnswerError(w, http.StatusBadRequest, "ErrInp", message)
	} else {
		AnswerSuccess(w, http.StatusAccepted)
	}
}
