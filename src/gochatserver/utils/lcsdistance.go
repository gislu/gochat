package utils


//LCS算法
func LCS(source, target string) (int, string) {
	runesS := []rune(source)
	runesT := []rune(target)
	sLen := len(runesS)
	tLen := len(runesT)
	lengths := make([][]int, sLen+1)
	for i := 0; i <= sLen; i++ {
		lengths[i] = make([]int, tLen+1)
	}

	for i := 0; i < sLen; i++ {
		for j := 0; j <tLen; j++ {
			if runesS[i] == runesT[j] {
				lengths[i+1][j+1] = lengths[i][j] + 1
			} else if lengths[i+1][j] > lengths[i][j+1] {
				lengths[i+1][j+1] = lengths[i+1][j]
			} else {
				lengths[i+1][j+1] = lengths[i][j+1]
			}
		}
	}


	s := make([]rune, 0, lengths[sLen][tLen])
	for x, y := sLen, tLen; x != 0 && y != 0; {
		if lengths[x][y] == lengths[x-1][y] {
			x--
		} else if lengths[x][y] == lengths[x][y-1] {
			y--
		} else {
			s = append(s, runesS[x-1])
			x--
			y--
		}
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return len(s), string(s)
}