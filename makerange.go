package piscine

func MakeRange(min, max int) []int {
	// If min is greater than or equal to max, return nil slice
	if min >= max {
		return nil
	}

	// Calculate the length of the slice
	length := max - min

	// Create a slice with the appropriate length
	result := make([]int, length)

	// Fill the slice with values from min to max-1
	for i := 0; i < length; i++ {
		result[i] = min + i
	}

	return result
}
