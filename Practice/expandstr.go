/*Write a program that takes a string and displays it with exactly 
three spaces between each word, with no spaces nor tabs at neither the beginning 
nor the end.
The string will be followed by a newline ('\n').
A word, in this exercise, is a sequence of visible characters.
If the number of arguments is not 1, or if there are no word, the program 
displays nothing.*/

package piscine

import (
		"os"
		 "github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	started := false

	for _, char := range input {
		if char == ' ' || char == '\t' {
			if started {
				z01.PrintRune(' ')
				z01.PrintRune(' ')
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
