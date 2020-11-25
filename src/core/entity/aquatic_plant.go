package entity //import "github.com/florestario/core/entity"

// AquaticPlant Entity that represents aquatic plant
type AquaticPlant struct {
	ID       int
	Variety  string
	Specie   *Specie
	BioCycle BioCycle
}

// Specie Entity that represents species
type Specie struct {
	Name  string
	Genus *Genus
}

// Genus Entity that represents genus
type Genus struct {
	ID   int
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
