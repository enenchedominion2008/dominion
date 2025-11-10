package piscine

// ForEach applies function f to each element of slice a.
func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}
