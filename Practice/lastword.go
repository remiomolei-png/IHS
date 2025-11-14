
func LastWord(s string) string {
	word := ""

	for i := len(s) - 1; i >= 0; i-- {
		
		if s[i] == ' ' && word == "" {
			continue
		}
		
		if s[i] == ' ' {
			break
		}
		
		word = string(s[i]) + word
	}

	return word + "\n"
}
