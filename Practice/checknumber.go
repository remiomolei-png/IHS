package main

import "fmt"

func Checkintstr(str string) string {
	for _, char := range str{
		if char >= '0' && char <= '9'{
			return "Number in the string"
			
		}
	
	}
	return "String is okay"
}

func main (){
	fmt.Println(Checkintstr("ere-9"))
}