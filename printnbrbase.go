package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		printString("NV")
		return
	}

	baseLen := len(base)

	if nbr < 0 {
		z01.PrintRune('-')
		if nbr == -9223372036854775808 {
			printInBase(-(nbr / baseLen), base, baseLen)
			z01.PrintRune(rune(base[-(nbr % baseLen)]))
			return
		}
		nbr = -nbr
	}

	printInBase(nbr, base, baseLen)
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' {
			return false
		}
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func printInBase(n int, base string, baseLen int) {
	if n >= baseLen {
		printInBase(n/baseLen, base, baseLen)
	}
	z01.PrintRune(rune(base[n%baseLen]))
}

func printString(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}
