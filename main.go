package main

import (
	"fmt"
)

func isMutant(adn []string) (b bool) {
	adnByte := byteSlice(adn)
	n := len(adn)

	if n < 4 {
		return false
	}

	k := 0
	for i, thread := range adn {
		for j := range thread {
			k += checkPos(adnByte, i, j)
			fmt.Printf("%s -> [%d,%d] -> %d\n", string(adnByte[i][j]), i, j, k)
			if k > 1 {
				b = true
				break
			}
		}
		if k > 1 {
			b = true
			break
		}
	}

	fmt.Println("k =", k)
	fmt.Println(adnByte)
	return
}

func byteSlice(a []string) [][]byte {
	r := make([][]byte, 0, len(a))
	for _, v := range a {
		r = append(r, []byte(v))
	}

	return r
}

func checkPos(adnByte [][]byte, i, j int) (c int) {
	c += checkH(adnByte, i, j)
	c += checkV(adnByte, i, j)
	// c += checkD(adn, v, i, j)

	return
}

func checkH(adnByte [][]byte, i, j int) (c int) {
	base := adnByte[i][j]
	aux := 0
	n := len(adnByte)
	equals := make([][]int, 0, 5)

	// check right elements
	for r := j; r < n; r++ {
		if base != adnByte[i][r] || aux == 4 {
			break
		}
		aux++
		equals = append(equals, []int{i, r})
	}

	// check left elements
	for l := j; l >= 0; l-- {
		if base != adnByte[i][l] || aux == 4 {
			break
		}
		aux++
		equals = append(equals, []int{i, l})
	}

	if aux >= 4 {
		c = 1
		patch(adnByte, equals)
	}

	return
}

func checkV(adnByte [][]byte, i, j int) (c int) {
	base := adnByte[i][j]
	aux := 0
	n := len(adnByte)
	equals := make([][]int, 0, 5)

	// check down elements
	for d := i; d < n; d++ {
		if base != adnByte[d][j] || aux == 4 {
			break
		}
		aux++
		equals = append(equals, []int{d, j})
	}

	// check up elements
	for u := j; u >= 0; u-- {
		if base != adnByte[u][j] || aux == 4 {
			break
		}
		aux++
		equals = append(equals, []int{u, j})
	}

	if aux >= 4 {
		c = 1
		patch(adnByte, equals)
	}

	return
}

// func checkD(adnByte [][]byte, i, j int) (c int) {
// 	base := adnByte[i][j]
// 	aux := 0
// 	n := len(adnByte)
// 	equals := make([][]int, 0, 5)

// 	// check down elements
// 	for d := i; d < n; d++ {
// 		if base != adnByte[d][j] || aux == 4 {
// 			break
// 		}
// 		aux++
// 		equals = append(equals, []int{d, j})
// 	}

// 	// check up elements
// 	for u := j; u >= 0; u-- {
// 		if base != adnByte[u][j] || aux == 4 {
// 			break
// 		}
// 		aux++
// 		equals = append(equals, []int{u, j})
// 	}

// 	if aux >= 4 {
// 		c = 1
// 		patch(adnByte, equals)
// 	}

// 	return
// }

func patch(adnByte [][]byte, index [][]int) {
	for x, ind := range index {
		i := ind[0]
		j := ind[1]
		adnByte[i][j] = byte(x)
	}
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
