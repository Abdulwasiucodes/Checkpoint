package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		for i, wordStart := 0, 0; i <= len(arg); i++ {
			if i == len(arg) || arg[i] == ' ' {
				for j := wordStart; j < i-1; j++ {
					z01.PrintRune(toLower(rune(arg[j])))
				}
				if wordStart < i {
					z01.PrintRune(toUpper(rune(arg[i-1])))
				}
				if i < len(arg) {
					z01.PrintRune(' ')
				}
				wordStart = i + 1
			}
		}
		z01.PrintRune('\n')
	}
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}

func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - ('a' - 'A')
	}
	return r
}

