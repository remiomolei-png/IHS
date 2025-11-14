package main

import "fmt"

func Capitalize(s string) string {
	runes := []rune(s)
	word := true

	for i:=0; i < len(runes); i++ {
		if (runes[i] >='a' && runes[i]<= 'z') || (runes[i] >='A' && runes[i]<= 'Z') || (runes[i] >='0' && runes[i]<= '9'){
			if word && runes[i]	>='a' && runes[i]<= 'z' {
				runes[i] -= 32
			} else if !word && runes[i] >='A' && runes[i]<= 'Z' {
				runes[i] += 32
			}
			word = false
		} else {
			word = true
		}

	}
	return string(runes)
}


func main() {
	fmt.Println(Capitalize("hello! how are you?"))
	fmt.Println(Capitalize("HEEEllo How Are You"))
	fmt.Println(Capitalize("Whats 4this 100K?"))
	fmt.Println(Capitalize("Whatsthis4"))
	fmt.Println(Capitalize("!!!!Whatsthis4"))
	fmt.Println(Capitalize(""))
}

