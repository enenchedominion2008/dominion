package piscine

func IterativePower(nb int, power int) int {
	// Handle negative powers
	if power < 0 {
		return 0
	}

	// Any number to the power of 0 is 1
	if power == 0 {
		return 1
	}

	result := 1
	for i := 0; i < power; i++ {
		result *= nb
	}

	return result
}
