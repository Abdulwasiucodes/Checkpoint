package piscine

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	newWord := true

	for _, char := range s {
		if newWord {
			if char >= 'a' && char <= 'z' {
				return false
			}
			newWord = false
		}

		if char == ' ' {
			newWord = true
		}
	}

	return true
}
