package piscine

func TrimAtoi(s string) int {
	sign := 1
	result := 0
	started := false

	for _, char := range s {
		if char == '-' && !started {
			sign = -1
		} else if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
			started = true
		} else if char == ' ' || char == '+' {
			started = false
		}
	}

	return result * sign
}
