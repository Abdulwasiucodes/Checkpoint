package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1, s2 := os.Args[1], os.Args[2]
	charMap := make(map[rune]bool)
	seen := make(map[rune]bool)

	for _, r := range s2 {
		charMap[r] = true
	}

	for _, r := range s1 {
		if charMap[r] && !seen[r] {
			z01.PrintRune(r)
			seen[r] = true
		}
	}

	z01.PrintRune('\n')
}
