package models

import (
	"fmt"
	"log"
)

func GetCurrencies() {
	rows, err := db.Query("select currency_code from currencies")
	if err != nil {
		log.Fatal(err)
		fmt.Printf("\n")
	}

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
