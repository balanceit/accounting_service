package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/balanceit/accounting_service/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		models.GetCurrencies()
		JSONResponse(w, "gotten", http.StatusOK)
	}
}

// JSONResponse attempts to set the status code, c, and marshal the given interface, d, into a response that
// is written to the given ResponseWriter.
func JSONResponse(w http.ResponseWriter, d interface{}, c int) {
	dj, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}
