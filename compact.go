package piscine

func Compact(ptr *[]string) int {
	originalSlice := *ptr
	nonZeroCount := 0
	j := 0

	for i := 0; i < len(originalSlice); i++ {
		if originalSlice[i] != "" {
			originalSlice[j] = originalSlice[i]
			j++
			nonZeroCount++
		}
	}

	*ptr = originalSlice[:j]
	return nonZeroCount
}
