package piscine
func PrintIf(str string) string {
    if len(str) >= 3 {
        return "G\n"
    } else if len(str) == 0 {
        return "G\n"
    } else {
        return "Invalid Input\n"
    }
}