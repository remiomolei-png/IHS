package main

func FirstWord(s string) string {
	word := ""
	started := false

	for _, char := range s{
		if char == ' ' && !started {
			continue
		}
		if char == ' '{
			break
		}
		started = true
		word += string(char)
	}
	return word + "\n"
}
