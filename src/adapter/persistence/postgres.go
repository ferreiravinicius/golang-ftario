package persistence

import (
	"database/sql"
	"errors"
	"github.com/florestario/core/entity"
	"github.com/florestario/core/gateway"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"sync"
)

var (
	once     sync.Once
	instance *PostgresShowroomPersistence
)

const (
	driver     = "pgx"
	dataSource = "user=florestario dbname=florestario password=florestario sslmode=disable"
)

type PostgresShowroomPersistence struct {
	db *sqlx.DB
}

func Instance() gateway.ShowroomPersistence {
	once.Do(func() {
		//TODO: tratar erro de conexão
		db, _ := sqlx.Connect(driver, dataSource)
		instance = &PostgresShowroomPersistence{db: db}
	})
	return instance
}

func (persistence *PostgresShowroomPersistence) SaveGenus(genus *entity.Genus) (*entity.Genus, error) {
	var query = "INSERT INTO genus (name) VALUES ($1) RETURNING id"
	if err := persistence.db.Get(&genus.ID, query, genus.Name); err != nil {
		return nil, err
	}
	return genus, nil
}

type GenusSchema struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func (persistence *PostgresShowroomPersistence) GetGenusByName(name string) (*entity.Genus, error) {
	const query = "SELECT id, name FROM genus WHERE UPPER(name) = UPPER($1) LIMIT 1"
	var existingGenus GenusSchema
	if err := persistence.db.Get(&existingGenus, query, name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &entity.Genus{
		ID:   existingGenus.ID,
		Name: existingGenus.Name,
	}, nil
}

func (persistence *PostgresShowroomPersistence) FilterGenusByName(name string) ([]entity.Genus, error) {
	const query = "SELECT id, name FROM genus WHERE UPPER(name) LIKE UPPER($1)"
	retrieved := make([]entity.Genus, 0)
	paramName := name + "%"
	if err := persistence.db.Select(&retrieved, query, paramName); err != nil {
		return nil, err
	}
	return retrieved, nil
}

func (persistence *PostgresShowroomPersistence) SavePlant(plant *entity.AquaticPlant) error {
	panic("implement me")
}

func (persistence *PostgresShowroomPersistence) FilterSpecieByName(name string) ([]entity.Specie, error) {
	const query = `SELECT s.id AS "id", s.name AS "name", 
					genus.id as "genus.id", genus.name as "genus.name"  
					FROM specie s 
					INNER JOIN genus ON s.genus_id = genus.id
					WHERE UPPER(s.name) LIKE UPPER($1)`

	retrieved := make([]entity.Specie, 0)
	paramName := name + "%"
	if err := persistence.db.Select(&retrieved, query, paramName); err != nil {
		return nil, err
	}
	return retrieved, nil
}

func (persistence *PostgresShowroomPersistence) FilterSpecieByNameAndGenus(name string, genusId int64) ([]entity.Specie, error) {
	const query = `SELECT s.id AS "id", s.name AS "name", 
					genus.id as "genus.id", genus.name as "genus.name"  
					FROM specie s 
					INNER JOIN genus ON s.genus_id = genus.id
					WHERE UPPER(s.name) LIKE UPPER($1)
					AND s.genus_id = $2`

	retrieved := make([]entity.Specie, 0)
	paramName := name + "%"
	if err := persistence.db.Select(&retrieved, query, paramName, genusId); err != nil {
		return nil, err
	}
	return retrieved, nil
}
