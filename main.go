package main

import (
	"fmt"
	"math/rand"
	"time"
)

var indexSet = map[string]struct{}{}

func isMutant(adn []string) (b bool) {
	n := len(adn)

	if n < 4 {
		return false
	}
	x := 1
	for len(indexSet) < n*n {
		i, j := randomIndex(n)
		base := string(adn[j][i])
		fmt.Printf("%d La base nitrogenada %s, se encuentra en la posición (%d,%d)\n", x, base, i, j)
		x++
	}

	return
}

func randomIndex(n int) (i, j int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	i = r1.Intn(n)
	j = r1.Intn(n)
	index := fmt.Sprintf("%d%d", i, j)
	if _, ok := indexSet[index]; !ok {
		indexSet[index] = struct{}{}
	}
	return
}

func main() {
	thread := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG",
	}
	_ = isMutant(thread)
}
