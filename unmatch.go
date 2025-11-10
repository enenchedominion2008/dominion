package piscine

func Unmatch(a []int) int {
	counts := make(map[int]int)

	for _, n := range a {
		counts[n]++
	}

	for k, v := range counts {
		if v%2 != 0 {
			return k
		}
	}
	return -1
}
