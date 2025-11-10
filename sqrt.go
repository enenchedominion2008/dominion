package piscine

func Sqrt(nb int) int {
	// Handle negative numbers and zero
	if nb <= 0 {
		return 0
	}

	// Check each number from 1 up to nb
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			return i
		}
		// If we've exceeded nb, no need to continue
		if i*i > nb {
			return 0
		}
	}
	return 0
}
