package piscine

import (
        "os"
        "github.com/01-edu/z01"
)

func main() {
    if len(os.Args) != 4 {
        z01.PrintRune('\n')
        return
    }

    str := os.Args[1]
    oldChar := os.Args[2]
    newChar := os.Args[3]

    if len(oldChar) != 1 || len(newChar) != 1{
        for _, char := range str {
            z01.PrintRune(char)
        }
        z01.PrintRune('\n')
        return
    }

    for _, char := range str {
        if char == rune(oldChar[0]) {
            z01.PrintRune(rune(newChar[0]))
        } else {
            z01.PrintRune(char)
        }
    }
    z01.PrintRune('\n')
}
