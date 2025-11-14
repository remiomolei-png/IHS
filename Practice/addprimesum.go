package main

import (
	"fmt"
	"os"
)

// Check if a number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	// Check odd divisors up to sqrt(n)
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Convert string to int manually (no imports)
func atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	
	result := 0
	start := 0
	negative := false
	
	// Check for negative sign
	if s[0] == '-' {
		negative = true
		start = 1
	} else if s[0] == '+' {
		start = 1
	}
	
	if start >= len(s) {
		return 0, false
	}
	
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		result = result*10 + int(s[i]-'0')
	}
	
	if negative {
		result = -result
	}
	
	return result, true
}

func main() {
	// Check if exactly 1 argument
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}
	
	// Convert argument to integer
	n, ok := atoi(os.Args[1])
	if !ok || n <= 0 {
		fmt.Println(0)
		return
	}
	
	// Calculate sum of primes
	sum := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	
	fmt.Println(sum)
}
