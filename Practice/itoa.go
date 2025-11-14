package main

import "fmt"

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	
	result := ""

	for n > 0 {
		nb := n%10
		result = string(nb + '0') + result 
		n /=10

	}
	return result
}

func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}