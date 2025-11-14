func findCommonChars(s1, s2 string) string {
	// Use a byte array for ASCII characters (more efficient than map for this use case)
	var seen [256]bool
	var result []byte

	// First pass: mark which characters exist in s2
	var inS2 [256]bool
	for i := 0; i < len(s2); i++ {
		inS2[s2[i]] = true
	}

	// Second pass: build result while maintaining order from s1
	for i := 0; i < len(s1); i++ {
		char := s1[i]
		if !seen[char] && inS2[char] {
			result = append(result, char)
			seen[char] = true
		}
	}

	return string(result)
}
