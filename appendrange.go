package piscine

func AppendRange(min, max int) []int {
	// If min is greater than or equal to max, return nil slice
	if min >= max {
		return nil
	}

	// Create an empty slice
	var result []int

	// Append values from min to max-1
	for i := min; i < max; i++ {
		result = append(result, i)
	}

	return result
}
