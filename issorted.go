package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	increasing := true
	decreasing := true

	for i := 0; i < len(a)-1; i++ {
		comparison := f(a[i], a[i+1])

		if comparison > 0 {
			increasing = false
		} else if comparison < 0 {
			decreasing = false
		}
	}

	return increasing || decreasing
}
