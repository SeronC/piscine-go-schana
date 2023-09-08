package piscine

func Itoa(num int) string {
	const digits = "0123456789"

	var result string
	negative := false

	if num < 0 {
		negative = true
		num = -num
	}

	for num > 0 {
		digit := num % 10
		result = string(digits[digit]) + result
		num /= 10
	}

	if negative {
		result = "-" + result
	} else if result == "" {
		result = "0"
	}

	return result
}
