package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/EvgenyiK/animals/server/datastore/animal"

	handlerAnimal "github.com/EvgenyiK/animals/server/delivery/animal"
	"github.com/EvgenyiK/animals/server/driver"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := driver.MySQLConfig{
		Host:     os.Getenv("HOST"),
		User:     os.Getenv("USERD"),
		Password: os.Getenv("PASSWORD"),
		Port:     os.Getenv("PORT"),
		Db:       os.Getenv("DB"),
	}
	var err error
	db, err := driver.ConnectToMySQL(conf)
	if err != nil {
		log.Println("could not connect to sql, err:", err)
		return
	}
	datastore := animal.New(db)
	handler := handlerAnimal.New(datastore)

	http.HandleFunc("/animal", handler.Handler)
	fmt.Println(http.ListenAndServe(":9000", nil))
}
