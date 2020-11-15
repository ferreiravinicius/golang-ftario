package entity

import (
	"github.com/florestario/core/err"
	"strings"
)

// AquaticPlant Entity that represents aquatic plant
type AquaticPlant struct {
	ID int
	Variety  string
	Specie   *Specie
	BioCycle BioCycle
}

// UniqueCode Generates a unique code based on Genus, Specie and Variety
func (plant AquaticPlant) UniqueCode() (string, error) {

	generateCode := func(text string) string {
		size := len(text)
		if size > 4 {
			size = 4
		}
		return text[0:size]
	}

	specie := plant.Specie
	if specie == nil || len(specie.Name) == 0 {
		return "", err.ErrWrongSpecie
	}

	genus := specie.Genus
	if genus == nil || len(genus.Name) == 0 {
		return "", err.ErrWrongGenus
	}

	varietyCode := generateCode(plant.Variety)
	specieCode := generateCode(specie.Name)
	genusCode := generateCode(plant.Specie.Genus.Name)
	code := genusCode + specieCode + varietyCode

	return strings.ToUpper(code), nil
}

// Specie Entity that represents species
type Specie struct {
	Name  string
	Genus *Genus
}

// Genus Entity that represents genus
type Genus struct {
	Name string
}

// BioCycle Biologic Life Cycle 
type BioCycle string

// Available Biologic Life Cycle 
const (
	Annual    BioCycle = "ANNUAL"
	Biennial  BioCycle = "BIENNIAL"
	Perennial BioCycle = "PERENNIAL"
)
