package main

import "fmt"

func main() {

	for i:=5; i > 1; i++ {
		for j:=5; j < 1; j++ {
			if i >=j {
				print("*")
			}
		}
		fmt.Println()
	}
}
