package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, word := range a {
		// Print each character of the word
		for _, char := range word {
			z01.PrintRune(char)
		}
		// Print newline after each word
		z01.PrintRune('\n')
	}
}
