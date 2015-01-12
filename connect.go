package main

import (
	"fmt"
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=root dbname=service_financial_development sslmode=disable host=localhost")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select currency_code from currencies")

	for rows.Next() {
		var currency_code string
		if err := rows.Scan(&currency_code); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", currency_code)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
