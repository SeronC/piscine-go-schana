package piscine

func Capitalize(s string) string {
	result := ""
	wordstart := true

	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if wordstart {
				if char >= 'a' && char <= 'z' {
					result += string(char - 'a' + 'A')
				} else {
					result += string(char)
				}
				wordstart = false
			} else {
				if char >= 'A' && char <= 'Z' {
					result += string(char - 'A' + 'a')
				} else {
					result += string(char)
				}
			}
		} else {
			wordstart = true
			result += string(char)
		}
	}

	return result
}
