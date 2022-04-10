package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type jsonExpect struct {
	Count_mutant_dna int     `json:"count_mutant_dna"`
	Count_human_dna  int     `json:"count_human_dna"`
	Ratio            float32 `json:"ratio"`
}

type postCase struct {
	jsonByte   []byte
	statusCode int
}

var postCases = []postCase{
	{[]byte(`{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`), 200},
	{[]byte(`{"dna":["ATGCGA","CAGTGC","TTATTT","AGACGG","GCGTCA","TCACTG"]}`), 403},
	{[]byte(`{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"],""}`), 400},
	{[]byte(`{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"],}`), 400},
	{[]byte(`{"adn":["ATGCGA","CAGTGC","TTATTT","AGACGG","GCGTCA","TCACTG"]}`), 400},
	{[]byte(`{"dna":["ATGCGA","CAGTGC","TTATTT","AGACGG","GCGTCA","TCXCTG"]}`), 400},
	{[]byte(`{"dna":["TTT","AAA","CCC"]}`), 400},
	{[]byte(`{"dna":"ACTGGTCAAC"}`), 400},
	{[]byte(`{"dna":[]}`), 400},
	{[]byte(`{"dna":""}`), 400},
}

// TestGetStats tests the estructure of the response for the
// endpoint '/stats'
func TestGetStats(t *testing.T) {
	req, err := http.NewRequest("GET", "/stats", nil)
	if err != nil {
		t.Fatalf("could not created get request: %s", err.Error())
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statistics)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var expected jsonExpect
	dec := json.NewDecoder(rr.Body)
	dec.DisallowUnknownFields()
	err = dec.Decode(&expected)

	if err != nil {
		t.Error(err.Error())
	}
}

// TestCheckMutant tests diferents JSON body for the endpoint '/mutants/'
// and compares the expected status code
func TestCheckMutant(t *testing.T) {
	for _, jsonBody := range postCases {
		req, err := http.NewRequest("POST", "/mutant/", bytes.NewBuffer(jsonBody.jsonByte))
		if err != nil {
			t.Fatalf("could not created post request: %s", err.Error())
		}

		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(checkMutant)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != jsonBody.statusCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, jsonBody.statusCode)
		}
	}
}
