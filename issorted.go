package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}
	asc := true
	desc := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			asc = false
		}
		if f(a[i], a[i+1]) < 0 {
			desc = false
		}
	}
	return asc || desc
}
