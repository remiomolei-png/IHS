package main

import (
	"fmt"
	"os"
)

// Convert string to int manually
func atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	
	result := 0
	
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		result = result*10 + int(s[i]-'0')
	}
	
	return result, true
}

func main() {
	// Check if exactly 1 argument
	if len(os.Args) != 2 {
		return
	}
	
	// Convert argument to integer
	n, ok := atoi(os.Args[1])
	if !ok || n <= 1 {
		return
	}
	
	// Find and print prime factors
	result := ""
	first := true
	
	// Start with 2 (smallest prime)
	for n%2 == 0 {
		if !first {
			result += "*"
		}
		first = false
		result += "2"
		n = n / 2
	}
	
	// Check odd factors starting from 3
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			if !first {
				result += "*"
			}
			first = false
			result += fmt.Sprintf("%d", i)
			n = n / i
		}
	}
	
	// If n is still greater than 1, then it's a prime factor itself
	if n > 1 {
		if !first {
			result += "*"
		}
		result += fmt.Sprintf("%d", n)
	}
	
	fmt.Println(result)
}
