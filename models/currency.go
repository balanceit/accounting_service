package models

import (
	"fmt"
	"log"
)

type Currency struct {
	Id           string `json:"id"`
	CurrencyCode string `json:"currency_code"`
	Rounding     string `json:"rounding"`
}

func GetCurrencies() []Currency {
	var rows, err = db.Query("select id, currency_code, rounding from currencies")

	if err != nil {
		log.Fatal(err)
		fmt.Printf("\n")
	}

	defer rows.Close()

	var currencies []Currency

	for rows.Next() {
		var c Currency
		if err := rows.Scan(&c.Id, &c.CurrencyCode, &c.Rounding); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", c)
		currencies = append(currencies, c)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return currencies
}
