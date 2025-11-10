package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Convert from baseFrom to decimal first
	decimalValue := 0
	baseFromLen := len(baseFrom)

	// Convert to decimal
	for i := 0; i < len(nbr); i++ {
		digit := string(nbr[i])
		digitValue := 0
		// Find the value of this digit in baseFrom
		for j := 0; j < len(baseFrom); j++ {
			if string(baseFrom[j]) == digit {
				digitValue = j
				break
			}
		}
		// Add to decimal value: digitValue * (base^position)
		power := len(nbr) - 1 - i
		decimalValue += digitValue * pow(baseFromLen, power)
	}

	// Convert from decimal to baseTo
	if decimalValue == 0 {
		return string(baseTo[0])
	}

	baseToLen := len(baseTo)
	result := ""

	for decimalValue > 0 {
		remainder := decimalValue % baseToLen
		result = string(baseTo[remainder]) + result
		decimalValue = decimalValue / baseToLen
	}

	return result
}

// Helper function to calculate power
func pow(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
