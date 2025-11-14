func isHidden(s1, s2 string) bool {
	i, j := 0, 0
	
	// Iterate through both strings
	for i < len(s1) && j < len(s2) {
		// If characters match, move to next character in s1
		if s1[i] == s2[j] {
			i++
		}
		// Always move to next character in s2
		j++
	}
	
	// If we've matched all characters in s1, it's hidden in s2
	return i == len(s1)
}
