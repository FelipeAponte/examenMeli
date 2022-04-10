package main

import (
	"encoding/json"
	"errors"
	"examenMeli/mutant"
	"log"
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

func checkMutant(w http.ResponseWriter, r *http.Request) {
	var thread dnaThread

	err := check.DecodeJSONBody(w, r, &thread)
	if err != nil {
		var mr *check.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		} else {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
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

func statistics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mutants := DB.QueryMutants()
	humans := DB.QueryHumans()
	ratio := float32(mutants) / float32(humans)

	stat := statistic{
		Count_mutant_dna: mutants,
		Count_human_dna:  humans,
		Ratio:            ratio,
	}

	json.NewEncoder(w).Encode(stat)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/mutant/", checkMutant).Methods("POST")
	router.HandleFunc("/stats", statistics).Methods("GET")

	http.ListenAndServe(":8000", router)
}
