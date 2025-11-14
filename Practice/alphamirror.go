/*Write a program called alphamirror that takes a string as 
argument and displays this string after replacing each alphabetical character with the opposite alphabetical character. 
The case of the letter remains unchanged, for example : 'a' 
becomes 'z', 'Z' becomes 'A' 'd' becomes 'w', 'M' becomes 'N' 
The final result will be followed by a newline ('\n'). 
If the number of arguments is different from 1, the program 
prints a new line. */
package main
import(
		"os"


		"github.com/01-edu/z01"
)
func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}


	input := os.Args[1]

	for _, char :=range input {
		if char >='a' && char <= 'z'{
			z01.PrintRune('z' - (char-'a'))
		} else if char >= 'A' && char <= 'Z'{
			z01.PrintRune('Z'- (char - 'A'))
		} else {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}