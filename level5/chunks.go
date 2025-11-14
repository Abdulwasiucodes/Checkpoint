package piscine

import "fmt"

func Chunk(slice []int, size int) {
	if size <= 0 {
		fmt.Println()
		return
	}
	
	var chunks [][]int
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}

	fmt.Println(chunks)
}
