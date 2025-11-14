//theres a mistake

package piscine
func CamelToSnakeCase(s string) string {
    // Check if the input string is empty
    if len(s) == 0 {
        return ""
    }

    var result []rune
    for i, c := range s {
        if i > 0 && c >= 'A' && c <= 'Z' {
            result = append(result, '_')
            result = append(result, c + 'a' - 'A')
        } else {
            result = append(result, c)
        }
    }
    return string(result)
}