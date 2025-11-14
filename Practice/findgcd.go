package main

import "fmt"

func findGCD(a, b int) int {
    for b!=0 {
        a,b = b, a%b
    }
    return a
}

func main() {
    a:= 50
    b:=10
    result := findGCD(a,b)

    fmt.Println(result)
}

