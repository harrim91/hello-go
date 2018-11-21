package arrmap

// Map applies the given transformation to every element in an array and returns the resulting array
func Map(fn func(int) int, arr []int) []int {
	res := make([]int, len(arr))

	for i, val := range arr {
		res[i] = fn(val)
	}

	return res
}
