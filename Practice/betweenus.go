/*Write a function named BetweenUs that takes 3 paramters and return :
If the first paramter is between the second and third paramters, return true else return false
If the second parameter is bigger than the third return false*/
package main

import "fmt"

func BetweenUs(num, min, max int) bool {
	if min > max {
		return false
 	}
	return num > min && num < max

}

func main(){
	fmt.Println(BetweenUs(1, 2, 3))
	fmt.Println(BetweenUs(4, 3, 5))
	fmt.Println(BetweenUs(8, 7, 9))
	fmt.Println(BetweenUs(1, 1, 1))
	fmt.Println(BetweenUs(1, 2, 1))
	fmt.Println(BetweenUs(-1, -10, 0))
}
