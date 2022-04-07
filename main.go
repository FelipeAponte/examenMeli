package main

import (
	"examenMeli/mutant"
	"fmt"
)

func main() {
	// thread := []string{
	// 	"ATGCGA",
	// 	"CAGTGC",
	// 	"TTATGT",
	// 	"AGAAGG",
	// 	"CCCCTA",
	// 	"TCACTG",
	// }

	thread := []string{
		"TGCTCTCGAT",
		"CACCTCGAGG",
		"ACAGTGCCAG",
		"GTCTGCCGCA",
		"CGATGAAGCC",
		"CTCATGCTAC",
		"TTCAGGTACA",
		"TATTCCGCAT",
		"AACGTAACGA",
		"CTTGATTTAC",
	}

	b := mutant.IsMutant(thread)
	fmt.Println("ES MUTANTE: ", b)
}
