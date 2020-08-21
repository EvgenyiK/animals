package datastore

import "animals/server/entities"

//Animal ...
type Animal interface{
	Get(id int)([]entities.Animal, error)
	Create(entities.Animal)(entities.Animal,error)
}