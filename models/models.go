package models

import (
	"fmt"
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB // global variable to share it between main and the HTTP handler

func Setup() error {
	var err error
	db, err = sql.Open("postgres", "user=root dbname=service_financial_development sslmode=disable host=localhost")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error on opening database connection: %s", err.Error())
		fmt.Printf("\n")
	}

	return nil
}
