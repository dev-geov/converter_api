package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dev-geov/converter_api/controllers"
	"github.com/dev-geov/converter_api/models"
	"github.com/gorilla/mux"
)

func Exchange(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	amount, _ := strconv.ParseFloat(params["amount"], 64)
	rate, _ := strconv.ParseFloat(params["rate"], 64)

	from := strings.ToUpper(params["from"])
	to := strings.ToUpper(params["to"])

	conversion := models.Conversion{Amount: amount, From: from, To: to, Rate: rate}
	models.DB.Create(&conversion)

	var symbol string

	value := amount * rate

	switch to {
	case "USD":
		symbol = "$"
	case "EUR":
		symbol = "£"
	case "BTC":
		symbol = "B"
	case "BRL":
		symbol = "R$"
	default:
		symbol = "#"
	}

	result := models.Result{Value: value, CoinSymbol: symbol, ConversionID: conversion.ID}
	models.DB.Create(&result)

	result_conversion := make(map[string]interface{})
	result_conversion["valorConvertido"] = result.Value
	result_conversion["simboloMoeda"] = result.CoinSymbol

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result_conversion)
}

func Routers() {
	router := mux.NewRouter()

	router.HandleFunc("/exchange/{amount}/{from}/{to}/{rate}", Exchange)
	router.HandleFunc("/exchange/conversions", controllers.GetConversions)
	router.HandleFunc("/exchange/results", controllers.GetResults)

	log.Println("Running on port 8080")
	http.ListenAndServe(":8080", router)
}

func main() {
	models.InitialConfig()
	Routers()
}
