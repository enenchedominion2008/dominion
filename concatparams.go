package piscine

func ConcatParams(args []string) string {
	// If the slice is empty, return empty string
	if len(args) == 0 {
		return ""
	}

	// Initialize result with the first element
	result := args[0]

	// Append remaining elements with newline separator
	for i := 1; i < len(args); i++ {
		result += "\n" + args[i]
	}

	return result
}
