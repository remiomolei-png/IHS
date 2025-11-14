package main

import "fmt"

func Checkstrlenghth(str string) string {
	if len(str) >= 3 || len(str)< 3{
		return "G"
	}
	return "Invalid Input"
	}

func main (){
	fmt.Println(Checkstrlenghth("amjjo"))
}