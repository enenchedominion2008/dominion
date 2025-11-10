package piscine

func Split(s, sep string) []string {
	var result []string
	start := 0

	for i := 0; i <= len(s)-len(sep); i++ {
		// Check if we found the separator at current position
		if s[i:i+len(sep)] == sep {
			// Add the substring from start to current position
			if i > start {
				result = append(result, s[start:i])
			}
			// Move start position past the separator
			start = i + len(sep)
			// Skip ahead since we found the separator
			i = start - 1
		}
	}

	// Add the remaining part after the last separator
	if start < len(s) {
		result = append(result, s[start:])
	}

	return result
}
