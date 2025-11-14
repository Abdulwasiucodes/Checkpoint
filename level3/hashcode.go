package piscine
func HashCode(str string) string {
    result := ""
    for i := 0; i < len(str); i++ {
        char := str[i]
        hashValue := (int(char) + len(str)) % 127
        if hashValue < 33 {
            hashValue += 33
        }
        result += string(hashValue)
    }
    return result
}