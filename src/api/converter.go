package api

import "github.com/florestario/core/entity"

type AquaticPlantInput struct {
	ID      int    `json:"id"`
	Variety string `json:"variety"`
	Specie  string `json:"specie"`
}

func (input *AquaticPlantInput) ToEntity() *entity.AquaticPlant {
	return &entity.AquaticPlant{
		ID:      input.ID,
		Specie:  &entity.Specie{Name: input.Specie},
		Variety: input.Variety,
	}
}

type AquaticPlantOutput struct {
	ID      int    `json:"id"`
	Variety string `json:"variety"`
	Specie  string `json:"specie"`
}

func FromEntity(entity *entity.AquaticPlant) *AquaticPlantOutput {

	var specie string
	if entity.Specie != nil {
		specie = entity.Specie.Name
	}

	return &AquaticPlantOutput{
		ID:      entity.ID,
		Variety: entity.Variety,
		Specie:  specie,
	}
}


