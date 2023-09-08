package piscine

func Compare(a, b string) int {
	if a == b {
		return 0 // Strings are equal
	}

	minLength := len(a)
	if len(b) < minLength {
		minLength = len(b)
	}

	for i := 0; i < minLength; i++ {
		if a[i] < b[i] {
			return -1 // a is lexicographically smaller
		} else if a[i] > b[i] {
			return 1 // b is lexicographically smaller
		}
	}

	// All common characters are equal up to minLength
	if len(a) < len(b) {
		return -1 // a is shorter, so it's lexicographically smaller
	} else {
		return 1 // b is shorter, so it's lexicographically smaller
	}
}
