package piscine

func isWhiteSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

func SplitWhiteSpaces(str string) []string {
	var tab []string
	start := 0
	inWord := false

	for i := 0; i < len(str); i++ {
		if isWhiteSpace(str[i]) {
			if inWord {
				tab = append(tab, str[start:i])
				inWord = false
			}
		} else if !inWord {
			start = i
			inWord = true
		}
	}

	if inWord {
		tab = append(tab, str[start:])
	}

	return tab
}
