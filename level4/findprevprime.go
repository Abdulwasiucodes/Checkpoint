package piscine
func FindPrevPrime(nb int) int {
// Start from the given number and decrement until a prime is found
for i := nb; i >= 2; i-- {
if isPrime(i) {
return i
}
}
// If no prime number is found, return 0
return 0
}
// Helper function to check if a number is prime
func isPrime(n int) bool {
if n < 2 {
return false
}
for i := 2; i*i <= n; i++ {
if n%i == 0 {
return false
}
}
return true
}