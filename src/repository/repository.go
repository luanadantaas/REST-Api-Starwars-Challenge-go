package repository

import "starwars/src/entity"

type Repository interface {
	AddPlanet(entity.Planet) error
	ListPlanet() ([]entity.Planet, error)
	FindById(id int) (*entity.Planet, error)
	FindByName(name string) (*entity.Planet, error)
	RemovePlanet(id int) error
}
