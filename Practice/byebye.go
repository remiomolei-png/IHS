/*Write a function that takes a slice of string's 
and returns a new slice without the first element.
If the slice is empty, return the empty slice.*/

package main

import "fmt"

func ByeByeFirst(strings []string) []string {
	if len(strings) == 0 {
		return strings
	}
	return strings[1:]
}

func main() {
	fmt.Println(ByeByeFirst([]string{}))
	fmt.Println(ByeByeFirst([]string{"one arg"}))
	fmt.Println(ByeByeFirst([]string{"first", "second"}))
	fmt.Println(ByeByeFirst([]string{"", "abcd", "efg"}))
}

	 

