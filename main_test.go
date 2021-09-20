package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_exchange_conversion(t *testing.T) {
	req, err := http.NewRequest("GET", "/exchange/20/brl/usd/5", nil)

	if err != nil {
		t.Fatal("error")
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Exchange)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
