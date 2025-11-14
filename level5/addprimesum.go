package main

import (
	"os"
	"github.com/01-edu/z01"
)

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func AddPrimeSum(n int) int {
	sum := 0
	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			sum += i
		}
	}
	return sum
}

func PrintInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	
	digits := make([]rune, 0)
	for n > 0 {
		digits = append([]rune{rune(n%10) + '0'}, digits...)
		n /= 10
	}
	
	for _, digit := range digits {
		z01.PrintRune(digit)
	}
}

func main() {
	if len(os.Args) != 2 || os.Args[1][0] < '0' || os.Args[1][0] > '9' {
		PrintInt(0)
		z01.PrintRune('\n')
		return
	}
	
	n := 0
	for _, c := range os.Args[1] {
		if c < '0' || c > '9' {
			PrintInt(0)
			z01.PrintRune('\n')
			return
		}
		n = n*10 + int(c-'0')
	}
	
	if n <= 0 {
		PrintInt(0)
		z01.PrintRune('\n')
		return
	}
	
	result := AddPrimeSum(n)
	PrintInt(result)
	z01.PrintRune('\n')
}
