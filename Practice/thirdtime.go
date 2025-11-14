package main

import (
	"github.com/01-edu/z01"
)

func ThirdTimeIsACharm(str string) string {
	if len(str) < 3 {
		z01.PrintRune('\n')
	}

	for i :=2; i < len(str) ; i+=3 {
		z01.PrintRune(rune(str[i]))
	}
	z01.PrintRune('\n')
}
