package piscine

func RecursivePower(nb int, power int) int {
	// Handle negative powers
	if power < 0 {
		return 0
	}

	// Base case: any number to the power of 0 is 1
	if power == 0 {
		return 1
	}

	// Recursive case: nb^power = nb * nb^(power-1)
	return nb * RecursivePower(nb, power-1)
}
