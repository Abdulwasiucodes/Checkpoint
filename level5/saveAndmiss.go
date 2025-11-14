package piscine

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}

	var result string
	for i := 0; i < len(arg); i += num * 2 {
		for j := i; j < i+num; j++ {
			if j < len(arg) {
				result += string(arg[j])
			}
		}
	}

	return result
}
