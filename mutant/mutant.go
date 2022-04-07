package mutant

import "fmt"

func IsMutant(adn []string) (b bool) {
	adnByte := byteSlice(adn)
	n := len(adn)

	if n < 4 {
		return false
	}

	k := 0
	for i, thread := range adn {
		for j := range thread {
			k += checkPos(adnByte, i, j)
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

	return
}

func byteSlice(a []string) [][]byte {
	r := make([][]byte, 0, len(a))
	for _, v := range a {
		r = append(r, []byte(v))
	}

	return r
}

func patch(adnByte [][]byte, index [][]int) {
	for x, ind := range index {
		i := ind[0]
		j := ind[1]
		adnByte[i][j] = byte(x)
	}
}

func checkPos(adnByte [][]byte, i, j int) (c int) {
	c += checkHorizontal(adnByte, i, j)
	c += checkVertical(adnByte, i, j)
	c += checkDiagRigth(adnByte, i, j)
	c += checkDiagLeft(adnByte, i, j)

	return
}

func checkHorizontal(adnByte [][]byte, i, j int) (c int) {
	base := adnByte[i][j]
	aux := 0
	n := len(adnByte)
	equals := make([][]int, 0, 5)

	// check right elements
	for r := j + 1; r < n; r++ {
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
		fmt.Printf("Horizontales Repetidas %v\n", equals)
		patch(adnByte, equals)
	}

	return
}

func checkVertical(adnByte [][]byte, i, j int) (c int) {
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
	for u := i - 1; u >= 0; u-- {
		if base != adnByte[u][j] || aux == 4 {
			break
		}
		aux++
		equals = append(equals, []int{u, j})
	}

	if aux >= 4 {
		c = 1
		fmt.Printf("Verticales L Repetidas %v\n", equals)
		patch(adnByte, equals)
	}

	return
}

func checkDiagRigth(adnByte [][]byte, i, j int) (c int) {
	base := adnByte[i][j]
	aux := 0
	n := len(adnByte)
	equals := make([][]int, 0, 5)

	// check [R]igth-[d]own elements
	R := j
	for d := i; d < n; d++ {
		if R == n-1 {
			break
		}
		if base != adnByte[d][R] || aux == 4 {
			break
		}
		aux++
		equals = append(equals, []int{d, R})
		R++
	}

	if aux >= 4 {
		c = 1
		fmt.Printf("Diagonales R Repetidas %v\n", equals)
		patch(adnByte, equals)
	}

	return
}

func checkDiagLeft(adnByte [][]byte, i, j int) (c int) {
	base := adnByte[i][j]
	aux := 0
	n := len(adnByte)
	equals := make([][]int, 0, 5)

	// check [L]eft-[d]own elements
	L := j
	for d := i; d < n; d++ {
		if L < 0 {
			break
		}
		if base != adnByte[d][L] || aux == 4 {
			break
		}
		aux++
		equals = append(equals, []int{d, L})
		L--
	}

	if aux >= 4 {
		c = 1
		fmt.Printf("Diagonales L Repetidas %v\n", equals)
		patch(adnByte, equals)
	}

	return c
}
