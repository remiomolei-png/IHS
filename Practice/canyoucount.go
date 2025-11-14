/*Your program will receive some arguments. Count how many characters 
they have in total and print them.

If the number of arguments is invalid it should print 0.*/
package main

import ("os"
		"fmt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(0)
		return
	}

	count :=0
	for _, arg := range os.Args[1:]{
		count += len(arg)
	}
	fmt.Println(count)
}
