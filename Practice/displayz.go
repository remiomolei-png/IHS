package main

import "fmt"

func displayz(s string) {
	for _, char := range s {
		if char == 'z' || char == 'Z' {
			fmt.Println(string(char))
		return 
		}
	}
	fmt.Println("z")
}

func main(){ 
	displayz("Helloz")
	displayz("noze")
	displayz("Hzlloz")
	displayz("Hello")
 }