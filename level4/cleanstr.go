package main

import (
	"github.com/01-edu/z01"
	"os"
)

func CleanStr(str string) string {
	var cleanStr []rune
	inWord := false
	for _, c := range str {
		if c == ' ' {
			if inWord {
				cleanStr = append(cleanStr, ' ')
				inWord = false
			}
		} else {
			inWord = true
			cleanStr = append(cleanStr, c)
		}
	}
	
	// Remove leading and trailing spaces
	for len(cleanStr) > 0 && cleanStr[0] == ' ' {
		cleanStr = cleanStr[1:]
	}
	for len(cleanStr) > 0 && cleanStr[len(cleanStr)-1] == ' ' {
		cleanStr = cleanStr[:len(cleanStr)-1]
	}

	return string(cleanStr)
}

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	cleanStr := CleanStr(os.Args[1])
	for _, c := range cleanStr {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}


