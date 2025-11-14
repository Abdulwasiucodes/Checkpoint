package piscine

import (
	"github.com/01-edu/z01"
)

func intToString(n int) string {
	if n == 0 {
		return "0"
	}

	var digits []rune
	isNegative := n < 0
	if isNegative {
		n = -n
	}

	for n > 0 {
		digits = append(digits, rune(n%10)+'0')
		n /= 10
	}

	if isNegative {
		digits = append(digits, '-')
	}

	// Reverse the digits slice
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return string(digits)
}

func ReduceInt(a []int, f func(int, int) int) {
	x := a[0]
	for i := 1; i < len(a); i++ {
		x = f(x, a[i])
	}

	n := intToString(x)
	for _, r := range n {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}