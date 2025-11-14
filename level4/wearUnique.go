package piscine

import "strings"

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	uc := make(map[rune]bool)

    for _,char:= range str1{
        if !strings.Contains(str2 , string(char)) && !uc[char]{
            uc[char] = true
        }
    }

    for _,char:= range str2{
        if !strings.Contains(str1 , string(char)) && !uc[char]{
            uc[char] = true
}
    }
    return len(uc)
}