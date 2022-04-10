// package main
package main

import (
	"encoding/json"
	"errors"
	"examenMeli/mutant"
	"net/http"

	check "examenMeli/checkJsonBody"
	DB "examenMeli/database"

	"github.com/gorilla/mux"
)

type dnaThread struct {
	Dna []string `json:"dna"`
}

type statistic struct {
	Count_mutant_dna int     `json:"count_mutant_dna"`
	Count_human_dna  int     `json:"count_human_dna"`
	Ratio            float32 `json:"ratio"`
}

// checkMutant is a Handler function to implement the method POST
// for the endpoint '/mutant/' and saves the correct data
func checkMutant(w http.ResponseWriter, r *http.Request) {
	var thread dnaThread

	err := check.DecodeJSONBody(w, r, &thread)
	if err != nil {
		var mr *check.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		}
		return
	}

	err = check.CheckThreadDna(thread.Dna)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	threadStr, _ := json.Marshal(thread.Dna)
	isMutant := mutant.IsMutant(thread.Dna)
	DB.DnaInsert(string(threadStr), isMutant)

	if isMutant {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// statistics is a Handler function to implement the method GET
// for the endpoint '/stats' and write a JSON response with the
// statistics of DNA strands
func statistics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mutants := DB.QueryMutants()
	humans := DB.QueryHumans()
	ratio := getRatio(mutants, humans)

	stat := statistic{
		Count_mutant_dna: mutants,
		Count_human_dna:  humans,
		Ratio:            ratio,
	}

	json.NewEncoder(w).Encode(stat)
}

// getRatio gets the ratio between mutants and humans
func getRatio(m, h int) (r float32) {
	if h == 0 {
		r = float32(m)
	} else {
		r = float32(m) / float32(h)
	}
	return
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/mutant/", checkMutant).Methods("POST")
	router.HandleFunc("/stats", statistics).Methods("GET")

	http.ListenAndServe(":8000", router)
}
