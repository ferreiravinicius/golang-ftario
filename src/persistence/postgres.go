package persistence

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/florestario/core/entity"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
)

type AquaticPostgres struct {
	db *sqlx.DB
}

var (
	once     sync.Once
	instance *AquaticPostgres
)

const (
	driver     = "pgx"
	dataSource = "user=florestario dbname=florestario password=florestario sslmode=disable"
)

func NewAquaticPostgres() *AquaticPostgres {
	once.Do(func() {
		//TODO: tratar erro de conexão
		db, _ := sqlx.Connect(driver, dataSource)
		instance = &AquaticPostgres{db: db}
	})
	return instance
}

func (pg *AquaticPostgres) SaveGenus(genus *entity.Genus) (*entity.Genus, error) {
	var query = "INSERT INTO genus (name) VALUES ($1) RETURNING id"
	pg.db.Get(&genus.ID, query, genus.Name)
	return genus, nil
}

// Return nil if not found
func (pg *AquaticPostgres) queryGenusByName(name string) (*GenusSchema, error) {
	const query = "SELECT id, name FROM genus WHERE name = $1 LIMIT 1"
	var existingGenus GenusSchema
	err := pg.db.Get(&existingGenus, query, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			//TODO: padronizar como retornar erro desconhecido
			return nil, fmt.Errorf("Unknown error when saving: %w ", err)
		}
	}
	return &existingGenus, nil
}

func (pg *AquaticPostgres) querySpecieByName(name string) (*SpecieSchema, error) {
	const query = "SELECT id, name FROM specie WHERE name = $1 LIMIT 1"
	var existingSpecie SpecieSchema
	err := pg.db.Get(&existingSpecie, query, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			//TODO: padronizar como retornar erro desconhecido
			return nil, fmt.Errorf("Unknown error when saving: %w ", err)
		}
	}
	return &existingSpecie, nil
}

func GetEmpty(text string) *string {
	if text == "" {
		return nil
	}
	return &text
}

func (pg *AquaticPostgres) SavePlant(plant *entity.AquaticPlant) error {

	specie := plant.Specie
	genus := specie.Genus

	tx := pg.db.MustBegin()

	existingGenus, err := pg.queryGenusByName(genus.Name)
	if err != nil {
		return err
	}

	var genusId int
	if existingGenus == nil {
		tx.Get(&genusId, "INSERT INTO genus (name) VALUES ($1) RETURNING id", genus.Name)
	} else {
		genusId = existingGenus.ID
	}

	existingSpecie, err := pg.querySpecieByName(specie.Name)
	if err != nil {
		return err
	}

	var specieId int
	if existingSpecie == nil {
		const querySpecie = "INSERT INTO specie (name, genus_id) VALUES ($1, $2) RETURNING id"
		tx.Get(&specieId, querySpecie, specie.Name, genusId)
	} else {
		specieId = existingSpecie.ID
	}

	const queryPlant = `
		INSERT INTO aquatic_plant (variety, specie_id)
		VALUES ($1, $2)
		RETURNING id
	`

	currentVariety := GetEmpty(plant.Variety)

	var gotId int
	tx.Get(&gotId, queryPlant, *currentVariety, specieId)
	tx.Commit()

	plant.ID = gotId

	return nil
}

func (pg *AquaticPostgres) GetPlant(specie *entity.Specie, variety string) (*entity.AquaticPlant, error) {
	return nil, nil
}

type GenusSchema struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type SpecieSchema struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type PlantSchema struct {
	ID      int          `db:"id"`
	Variety *string      `db:"variety"`
	Specie  SpecieSchema `db:"specie"`
}

func Other() {
	db, err := sqlx.Connect("pgx", "user=florestario dbname=florestario password=florestario sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	query := `
		SELECT p.id, p.variety, s.id "specie.id", s.name "specie.name"
		FROM aquatic_plant p
		INNER JOIN specie s ON p.specie_id = s.id
		WHERE p.id = 2
		LIMIT 1`

	//query := `SELECT p.id, p.variety FROM aquatic_plant p LIMIT 1`

	var result PlantSchema
	err = db.Get(&result, query)
	if err != nil {
		log.Fatalln(err)
	}
}

func Test() {
	db, err := sqlx.Connect("pgx", "user=florestario dbname=florestario password=florestario sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	name := "anotherName"

	tx := db.MustBegin()
	var id int
	tx.Get(&id, "INSERT INTO genus (name) VALUES ($1) RETURNING id", name)
	//tx.MustExec("INSERT INTO genus (name) VALUES ($1)", name)
	tx.Commit()

	//var id int
	//tx.Get(&id, "SELECT id FROM genus WHERE name = ?", name)

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}
