package main

func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	retrun 0
}