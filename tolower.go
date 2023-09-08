package piscine

func ToLower(s string) string {
	result := []rune(s)
	for i, char := range result {
		if char >= 'A' && char <= 'Z' {
			result[i] = char - 'A' + 'a'
		}
	}
	return string(result)
}
