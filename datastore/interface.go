package datastore

import "github.com/EvgenyiK/animals/entities"

type Animal interface {
	Get(id int) ([]entities.Animal, error)
	Create(entities.Animal) (entities.Animal, error)
}
