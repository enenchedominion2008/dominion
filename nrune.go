package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}

	index := 1
	for _, r := range s {
		if index == n {
			return r
		}
		index++
	}
	return 0
}
