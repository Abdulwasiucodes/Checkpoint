package main

import (
    "github.com/01-edu/z01"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        return
    }

    input := os.Args[1]
    var words [][]rune
    word := []rune{}

  for _, r := range input {
        if r != ' ' {
            word = append(word, r)
        } else {
            if len(word) > 0 {
                words = append(words, word)
                word = []rune{}
            }
        }
    }

    if len(word)>0{
        words = append(words, word)
    }
   

   for i, w := range words {
        for _, r := range w {
            z01.PrintRune(r)
        }
        if i < len(words)-1 {
            z01.PrintRune(' ')
            z01.PrintRune(' ')
            z01.PrintRune(' ')
        }
    }

    z01.PrintRune('\n')
}
