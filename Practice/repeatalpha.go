func RepeatAlpha(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char >= 'a' && char <= 'z' {
			
			index := int(char - 'a' + 1)
			
			for j := 0; j < index; j++ {
				result += string(char)
			}
		} else if char >= 'A' && char <= 'Z' {
			
			index := int(char - 'A' + 1)
			
			for j := 0; j < index; j++ {
				result += string(char)
			}
		} else {
			
			result += string(char)
		}
	}

	return result
}
