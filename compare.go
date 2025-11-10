package piscine

func Compare(a, b string) int {
	lenA := len(a)
	lenB := len(b)
	minLen := lenA
	if lenB < lenA {
		minLen = lenB
	}

	for i := 0; i < minLen; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}

	if lenA < lenB {
		return -1
	} else if lenA > lenB {
		return 1
	}
	return 0
}
