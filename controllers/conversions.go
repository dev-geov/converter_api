package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dev-geov/converter_api/models"
)

func GetConversions(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var conversions []models.Conversion
	log.Println(r.URL.Path, r.Method)
	w.Header().Set("Content-Type", "application/json")

	if query != nil {
		filters := make(map[string]interface{})
		for key, value := range query {
			filters[key] = value[0]
		}
		models.DB.Find(&conversions, filters)
	} else {
		models.DB.Find(&conversions)
	}
	json.NewEncoder(w).Encode(conversions)
}

func GetResults(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var results []models.Result
	log.Println(r.URL.Path, r.Method)
	w.Header().Set("Content-Type", "application/json")

	if query != nil {
		filters := make(map[string]interface{})
		for key, value := range query {
			filters[key] = value[0]
		}
		models.DB.Preload("Conversion").Find(&results, filters)
	} else {
		models.DB.Preload("Conversion").Find(&results)
	}
	json.NewEncoder(w).Encode(results)
}
