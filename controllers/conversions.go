package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dev-geov/converter_api/models"
)

func GetConversions(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, r.Method)
	w.Header().Set("Content-Type", "application/json")
	var conversions []models.Conversion
	models.DB.Find(&conversions)
	json.NewEncoder(w).Encode(conversions)
}

func GetResults(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, r.Method)
	w.Header().Set("Content-Type", "application/json")
	var results []models.Result
	models.DB.Preload("Conversion").Find(&results)
	json.NewEncoder(w).Encode(results)
}
