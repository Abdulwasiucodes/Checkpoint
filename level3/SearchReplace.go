//package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	// Check if the number of arguments is exactly 3
	if len(os.Args) != 4 {
		return
	}

	// Get the input string, the letter to be replaced, and the replacement letter
	inputStr := os.Args[1]
	letterToReplace := []rune(os.Args[2])[0]
	replacementLetter := []rune(os.Args[3])[0]

	// Iterate through the input string
	for _, char := range []rune(inputStr) {
		// If the current character matches the letter to be replaced, print the replacement letter
		if char == letterToReplace {
			z01.PrintRune(replacementLetter)
		} else {
			z01.PrintRune(char)
		}
	}

	// Print a newline
	z01.PrintRune('\n')
}