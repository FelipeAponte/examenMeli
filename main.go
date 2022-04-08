package main

import (
	"examenMeli/mutant"
	"fmt"
)

func main() {
	thread := []string{
		"GGCCCGGTGC",
		"TATATCGTTG",
		"TTTACTGTCT",
		"AGGCCGGCCT",
		"GTGATGTAAG",
		"ACGGCTCCAG",
		"CCTAGGCATG",
		"TGATGGACGA",
		"GATTGTAAAG",
		"TTCATGGAGT",
	}

	b := mutant.IsMutant(thread)
	fmt.Println("ES MUTANTE: ", b)
}
