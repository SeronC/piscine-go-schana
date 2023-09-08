package piscine

func IterativeFactorial(nb int) int {
	count := nb - 1
	number := nb
	if number >= 0 && number < 26 {
		if number == 0 {
			number = 1
		} else {
			for i := count; i > 0; i-- {
				number = number * i
			}
		}
	} else {
		number = 0
	}
	return number
}
