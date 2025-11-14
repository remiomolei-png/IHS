package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
		return
	}

	input := os.Args[1]
	started := false

	for _, char := range input {
		if char == ' ' || char == '\t'{
			if started {
				z01.PrintRune(' ')
				started = false
			}
		} else {
			z01.PrintRune(char)
			started = true
		}
	}
	z01.PrintRune('\n')	
}
