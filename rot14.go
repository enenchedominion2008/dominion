package piscine

func Rot14(s string) string {
	result := ""

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			newChar := ((char - 'a' + 14) % 26) + 'a'
			result += string(newChar)
		} else if char >= 'A' && char <= 'Z' {
			newChar := ((char - 'A' + 14) % 26) + 'A'
			result += string(newChar)
		} else {
			result += string(char)
		}
	}

	return result
}
