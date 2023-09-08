package piscine

func IterativePower(nb, power int) int {
	if power < 0 {
		return 0 // Power is not defined for negative exponents
	}

	result := 1
	for i := 0; i < power; i++ {
		result *= nb
	}
	return result
}
