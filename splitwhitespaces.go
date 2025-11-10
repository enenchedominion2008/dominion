package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	var currentWord string

	for _, char := range s {
		// Check if character is a whitespace (space, tab, or newline)
		if char == ' ' || char == '\t' || char == '\n' {
			// If we've built up a word, add it to the result
			if currentWord != "" {
				result = append(result, currentWord)
				currentWord = ""
			}
		} else {
			// Add character to current word
			currentWord += string(char)
		}
	}

	// Don't forget the last word if there is one
	if currentWord != "" {
		result = append(result, currentWord)
	}

	return result
}
