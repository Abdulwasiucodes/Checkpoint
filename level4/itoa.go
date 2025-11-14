package piscine

func Itoa(n int) string {
if n == 0 {
		return "0"
	}

	negative := false
	if n < 0 {
		negative = true
		n = -n
	}

	result := ""
	for n > 0 {
		digit := n % 10
		result = string('0'+digit) + result
		n /= 10
	}

	if negative {
		result = "-" + result
	}

	return result
}