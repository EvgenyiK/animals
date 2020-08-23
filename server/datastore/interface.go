package datastore

import "github.com/EvgenyiK/animals/server/entities"

//Animal ...
type Animal interface {
	Get(id int) ([]entities.Animal, error)
	Create(entities.Animal) (entities.Animal, error)
}
