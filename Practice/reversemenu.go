package main

import (
    "fmt"
)

func ReverseMenuIndex(menu []string) {
    length := len(menu)
    for i := 0; i < length/2; i++ {
        menu[i], menu[length-1-i] = menu[length-1-i], menu[i]
    }
}
