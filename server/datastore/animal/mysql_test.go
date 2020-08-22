package animal

import (
	"database/sql"
	//"log"
	"os"
	"reflect"

	//"github.com/joho/godotenv"

	"github.com/EvgenyiK/animals/server/entities"

	"github.com/EvgenyiK/animals/server/driver"

	"testing"
)



func initializeMySQL(t *testing.T) *sql.DB {
	conf := driver.MySQLConfig{
		Host:     os.Getenv("HOST"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Port:     os.Getenv("PORT"),
		Db:       os.Getenv("DB"),
	}

	var err error
	db, err := driver.ConnectToMySQL(conf)
	if err != nil {
		t.Errorf("could not connect to sql, err:%v", err)
	}

	return db
}

func TestDatastore(t *testing.T)  {
	db:=initializeMySQL(t)
    a:= New(db)
	testAnimalStorer_Get(t,a)
	testAnimalStorer_Create(t,a)
}


func testAnimalStorer_Create(t *testing.T, db AnimalStorer)  {
	testcases:= []struct{
		req entities.Animal
		response entities.Animal
	}{
		{entities.Animal{Name: "Hen", Age: 1},entities.Animal{3, "Hen",  1}},
		{entities.Animal{Name: "Pig", Age: 2},entities.Animal{4, "Pig",  2}},
	}
	for i, v := range testcases {
		resp,_:=db.Create(v.req)
		if !reflect.DeepEqual(resp, v.response) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.response)
		}
	}
}

func testAnimalStorer_Get(t *testing.T, db AnimalStorer)  {
	testcases:= []struct{
		id int
		resp []entities.Animal
	}{
		{0,[]entities.Animal{{1,"Hippo",10}, {2,"Ele", 20}}},
		{1,[]entities.Animal{{1,"Hippo",10}}},
	}
	for i, v := range testcases {
		resp,_:=db.Get(v.id)
		if !reflect.DeepEqual(resp, v.resp) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.resp)
		}
	}
}