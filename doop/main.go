package main

import "os"

func PrintRune(r rune) {
	os.Stdout.Write([]byte(string(r)))
}

func printNumber(n int64) {
	if n < 0 {
		PrintRune('-')
		n = -n
	}
	if n >= 10 {
		printNumber(n / 10)
	}
	PrintRune(rune(n%10 + '0'))
}

func isNumber(s string) (int64, bool) {
	var n int64
	if s == "" {
		return 0, false
	}
	neg := false
	i := 0
	if s[0] == '-' {
		neg = true
		i++
	} else if s[0] == '+' {
		i++
	}
	if i == len(s) {
		return 0, false
	}
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		n = n*10 + int64(s[i]-'0')
	}
	if neg {
		n = -n
	}
	return n, true
}

func safeAdd(a, b int64) (int64, bool) {
	if (b > 0 && a > 9223372036854775807-b) || (b < 0 && a < -9223372036854775808-b) {
		return 0, false
	}
	return a + b, true
}

func safeSub(a, b int64) (int64, bool) {
	if (b < 0 && a > 9223372036854775807+b) || (b > 0 && a < -9223372036854775808+b) {
		return 0, false
	}
	return a - b, true
}

func safeMul(a, b int64) (int64, bool) {
	if a == 0 || b == 0 {
		return 0, true
	}
	if a > 0 && b > 0 && a > 9223372036854775807/b {
		return 0, false
	}
	if a > 0 && b < 0 && b < -9223372036854775808/a {
		return 0, false
	}
	if a < 0 && b > 0 && a < -9223372036854775808/b {
		return 0, false
	}
	if a < 0 && b < 0 && a < 9223372036854775807/b {
		return 0, false
	}
	return a * b, true
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, okA := isNumber(os.Args[1])
	b, okB := isNumber(os.Args[3])
	op := os.Args[2]

	if !okA || !okB {
		return
	}

	switch op {
	case "+":
		if res, ok := safeAdd(a, b); ok {
			printNumber(res)
			PrintRune('\n')
		}
	case "-":
		if res, ok := safeSub(a, b); ok {
			printNumber(res)
			PrintRune('\n')
		}
	case "*":
		if res, ok := safeMul(a, b); ok {
			printNumber(res)
			PrintRune('\n')
		}
	case "/":
		if b == 0 {
			for _, c := range "No division by 0" {
				PrintRune(c)
			}
			PrintRune('\n')
		} else {
			printNumber(a / b)
			PrintRune('\n')
		}
	case "%":
		if b == 0 {
			for _, c := range "No modulo by 0" {
				PrintRune(c)
			}
			PrintRune('\n')
		} else {
			printNumber(a % b)
			PrintRune('\n')
		}
	}
}
