package piscine

func IsPrime(nb int) bool {
	// 1 and numbers less than 2 are not prime
	if nb <= 1 {
		return false
	}

	// 2 is the only even prime number
	if nb == 2 {
		return true
	}

	// Even numbers greater than 2 are not prime
	if nb%2 == 0 {
		return false
	}

	// Check odd divisors up to the square root
	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}

	return true
}
