package piscine

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0 // Assuming we're not dealing with fractional powers
	}

	return nb * RecursivePower(nb, power-1)
}
