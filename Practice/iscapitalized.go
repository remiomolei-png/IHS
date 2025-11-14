package piscine

func IsCapitalized(s string) bool {
	for _, char := range s {
		if char[0] >='A' || char[0] <= 'Z'{
			return true
		}
	}
	return false
}