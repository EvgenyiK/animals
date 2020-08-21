package animal

import (
	"github.com/EvgenyiK/animals/server/driver"
	"database/sql"
	"os"
	"testing"
)



func initializeMySQL(t *testing.T) *sql.DB {
	conf:= driver.MySQLConfig{
		Host: os.Getenv("COMPOSE_PROJECT_NAME-database"),
		User: os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Port: os.Getenv("HOST_MACHINE_MYSQL_PORT"),
		Db: os.Getenv("MYSQL_DATABASE"),
	}

	var err error
	db,err:= driver.ConnectToMySQL(conf)
	if err != nil {
		t.Errorf("could not connect to sql, err:%v", err)
	}

	return db
}

func TestDatastore(t *testing.T)  {
	db:=initializeMySQL(t)
    a:= New(db)
	//testAnimalStorer_Get(t,a)
	//testAnimalStorer_Create(t,a)
	return a
}