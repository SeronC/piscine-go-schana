package piscine

func NRune(str string, n int) rune {
	nrune := []rune(str)
	if n > len(str) {
		return 0
	}
	if n <= 0 {
		return 0
	}

	return nrune[n-1]
}
