package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}
	seen := make(map[rune]bool)
	combined := os.Args[1] + os.Args[2]
	for _, char := range combined {
		if !seen[char] {
			
			z01.PrintRune(char)
			seen[char] = true
		}
	}
	z01.PrintRune('\n')
}