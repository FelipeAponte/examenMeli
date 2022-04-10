//package mutant implements several functions that, given a
//DNA strand, return wherher a person is mutant or not
package mutant

// IsMutant returns true if a DNA strand has two contiguous
// sequences of four nitrogenous bases, or false otherwise
func IsMutant(adn []string) (b bool) {
	adnByte := byteSlice(adn)

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

// byteSlice converts a slice of strings to a slice of bytes
func byteSlice(a []string) [][]byte {
	n := len(a)
	r := make([][]byte, 0, n)

	for _, v := range a {
		r = append(r, []byte(v))
	}

	return r
}

// patch replaces bytes in a slice of bytes
// given several positions i, j
func patch(adnByte [][]byte, index [][]int) {
	for x, ind := range index {
		i := ind[0]
		j := ind[1]
		adnByte[i][j] = byte(x)
	}
}

// checkPos analyzes the neighborhood of a position [i, j] in
// a portion of bytes and returns the number of contiguous
// sequences found
func checkPos(adnByte [][]byte, i, j int) (c int) {
	c += checkHorizontal(adnByte, i, j)
	c += checkVertical(adnByte, i, j)
	c += checkDiagRight(adnByte, i, j)
	c += checkDiagLeft(adnByte, i, j)

	return
}

// checkHorizontal return 1 if found 4 equal bytes contigouos
// in the same horizontal line
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
		patch(adnByte, equals)
	}

	return
}

// checkVertical return 1 if found 4 equal bytes contigouos
// in the same vertical line
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
		patch(adnByte, equals)
	}

	return
}

// checkDiagRight return 1 if found 4 equal bytes contigouos
// in the same diagonal right line
func checkDiagRight(adnByte [][]byte, i, j int) (c int) {
	base := adnByte[i][j]
	aux := 0
	n := len(adnByte)
	equals := make([][]int, 0, 5)

	// check [R]ight-[d]own elements
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
		patch(adnByte, equals)
	}

	return
}

// checkDiagLeft return 1 if found 4 equal bytes contigouos
// in the same diagonal left line
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
		patch(adnByte, equals)
	}

	return
}
