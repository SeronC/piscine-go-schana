package piscine

func StrRev(s string) string {
	runeSlice := []rune(s)
	length := len(runeSlice)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
	}

	return string(runeSlice)
}
