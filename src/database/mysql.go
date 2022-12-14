package database

import (
	"database/sql"
	"fmt"
	"os"
	"starwars/src/entity"
	"starwars/src/logger"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUrl = "DATABASE"
)

type Repository struct {
	sqlDB *sql.DB
}

func New() (*Repository, error) {

	url := os.Getenv(dbUrl)
	path := fmt.Sprintf("root:password@tcp(" + url + ":3306)/task")
	db, err := sql.Open("mysql", path)

	if err != nil {
		logger.Log().Warn("couldn't open database: %v", err)
		return nil, err
	}

	return &Repository{sqlDB: db}, nil
}

func (d *Repository) Close() {
	d.sqlDB.Close()
}

func (d *Repository) AddPlanet(p entity.Planet) error {
	insert := "INSERT INTO planet (name, climate, terrain, appearances) VALUES ( ?, ?, ?)"
	_, err := d.sqlDB.Exec(insert, p.Name, p.Climate, p.Terrain, &p.Appearances)
	if err != nil {
		return err
	}

	return err
}

func (d *Repository) ListPlanet() ([]entity.Planet, error) {
	rows, err := d.sqlDB.Query("SELECT ID, name, climate, terrain, appearances FROM planet")
	if err != nil {
		logger.Log().Errorf("couldn't list data: %v", err)
		return nil, err
	}
	defer rows.Close()

	var planets []entity.Planet

	//goes through all the rows and adds the data in the slice
	for rows.Next() {
		var p entity.Planet
		if err := rows.Scan(&p.ID, &p.Name,
			&p.Climate, &p.Terrain, &p.Appearances); err != nil {
			return planets, err
		}
		planets = append(planets, p)
	}

	return planets, nil

}

func (d *Repository) FindById(id int) (*entity.Planet, error) {
	p := &entity.Planet{}
	find := "SELECT ID, name, climate, terrain, appearances FROM planet WHERE ID = ?"
	err := d.sqlDB.QueryRow(find, id).Scan(&p.ID, &p.Name,
		&p.Climate, &p.Terrain, &p.Appearances)
	if err != nil {
		logger.Log().Errorf("couldn't get data from planet %d: %v", id, err)
		return nil, err
	}

	return p, nil
}

func (d *Repository) FindByName(name string) (*entity.Planet, error) {
	p := &entity.Planet{}
	find := "SELECT ID, name, climate, terrain, appearances FROM planet WHERE name = ?"
	err := d.sqlDB.QueryRow(find, name).Scan(&p.ID, &p.Name,
		&p.Climate, &p.Terrain, &p.Appearances)
	if err != nil {
		logger.Log().Errorf("couldn't get data from planet %d: %v", name, err)
		return nil, err
	}

	return p, nil
}

func (d *Repository) RemovePlanet(id int) error {
	r := "DELETE FROM task WHERE name = ?"
	_, err := d.sqlDB.Exec(r, id)
	if err != nil {
		logger.Log().Errorf("coundn't remove planet %d: %v", id, err)
		return err
	}

	return err

}
