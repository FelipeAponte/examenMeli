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

	// Solo tinene una coincidencia usar en unit test
	// thread := []string{
	// 	"TGCTCTCGAT",
	// 	"CACCTCGAGG",
	// 	"ACAGTGCCAG",
	// 	"GTCTGCCGCA",
	// 	"CGATGAAGCC",
	// 	"CTCATGCTAC",
	// 	"TTCAGGTACA",
	// 	"TATTCCGCAT",
	// 	"AACGTAACGA",
	// 	"CTTGATTTAC",
	// }

	thread := []string{
		"CGGTATTGAC",
		"CAGCACAAGC",
		"GGTTTAGATA",
		"CCTATAGTCC",
		"ACGAGACTCT",
		"ATGTAGCCAA",
		"CCCCGACCTT",
		"CGGGTTGAAG",
		"CGTTCCCGGG",
		"CAAAACAGTA",
	}

	b := mutant.IsMutant(thread)
	fmt.Println("ES MUTANTE: ", b)
}
