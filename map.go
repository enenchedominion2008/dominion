package piscine

// Map applies the function f to each element of the int slice a
// and returns a slice of bool values containing the results.
func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, v := range a {
		result[i] = f(v)
	}
	return result
}
