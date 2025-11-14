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
	i := 0

	for _, r := range s2 {
		if i < len(s1) && r == rune(s1[i]) {
			i++
		}
	}

	if i == len(s1) {
		z01.PrintRune('1')
	} else {
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')
}
