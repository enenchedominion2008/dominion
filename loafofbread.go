package piscine

func LoafOfBread(str string) string {
	// If string empty or only spaces → just newline
	onlySpaces := true
	for _, r := range str {
		if r != ' ' {
			onlySpaces = false
			break
		}
	}
	if str == "" || onlySpaces {
		return "\n"
	}

	runes := []rune(str)

	// Count total non-space runes
	total := 0
	for _, r := range runes {
		if r != ' ' {
			total++
		}
	}
	if total < 5 {
		return "Invalid Output\n"
	}

	result := []rune{}
	count := 0

	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if r == ' ' {
			continue
		}

		result = append(result, r)
		count++

		if count == 5 {
			result = append(result, ' ')
			count = 0

			// skip exactly one character from the original string (including spaces)
			if i+1 < len(runes) {
				i++
			}
		}
	}

	// remove any trailing space
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return string(result) + "\n"
}
