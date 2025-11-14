package piscine
func RepeatAlpha(s string) string {
	var result string

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'a' && c <= 'z' {
			index := int(c - 'a' + 1)
			for j := 0; j < index; j++ {
				result += string(c)
			}
		} else if c >= 'A' && c <= 'Z' {
			index := int(c - 'A' + 1)
			for j := 0; j < index; j++ {
				result += string(c)
			}
		} else {
			result += string(c)
		}
	}

	return result
}