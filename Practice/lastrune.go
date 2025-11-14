package main

func LastRune(s string) rune {
	if len(s) == 0 {
		return 0
	}
	return []rune(s)[len([]rune(s))-1]
}
