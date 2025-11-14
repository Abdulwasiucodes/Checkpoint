package piscine
func Gcd(a, b uint) uint {
    if a == 0 {
        return b
    }
    if b == 0 {
        return a
    }
    for b != 0 {
        a, b = b, a%b
    }
    return a
}