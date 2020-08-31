package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"

	 handlerAnimal "github.com/EvgenyiK/animals/delivery/animal"
	"github.com/EvgenyiK/animals/datastore/animal"
	"github.com/EvgenyiK/animals/driver"
)


func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func main() {
	conf := driver.MySQLConfig{
		Host:     os.Getenv("SQL_HOST"),
		User:     os.Getenv("SQL_USER"),
		Password: os.Getenv("SQL_PASSWORD"),
		Port:     os.Getenv("SQL_PORT"),
		Db:       os.Getenv("SQL_DB"),
	}

	var err error

	db,err:= driver.ConnectToMySQL(conf)
	if err != nil {
		log.Println("could not connect to sql, err:", err)
		return
	}

	datastore:= animal.New(db)
	handler:= handlerAnimal.New(datastore)

	http.HandleFunc("/animal", handler.Handler)
	fmt.Println(http.ListenAndServe(":9000", nil))
}
