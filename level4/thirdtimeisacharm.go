package piscine

func ThirdTimeIsACharm(str string) string {
	if len(str) < 3 {
		return "\n"
	}
	var result string
	for i, r := range str {
		if (i+1)%3 == 0 {
			result += string(r)
		}
	}
	return result + "\n"
}