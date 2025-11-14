package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 1 {
		return
	}

	first := true
	for d := 2; num > 1; d++ {
		for num%d == 0 {
			if !first {
				fmt.Print("*")
			}
			fmt.Print(d)
			num /= d
			first = false
		}
	}
	if first {
		fmt.Print(os.Args[1])
	}
	fmt.Println()
}