package piscine

func Rot13(s string) string {
	rot := make([]rune, len(s))
	for i, char := range s {
		if char >= 'a' && char <= 'z' {
			rot[i] = 'a' + (char - 'a' + 13) %26
		} else if char >= 'A' && char <= 'Z'{
			rot[i] = 'A' + (char - 'A' + 13) %26
		} else {
			rot[i] = char1
		}
	}
	return string rot
}