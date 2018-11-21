package fibonacci

// Fibonacci returns an array of length i containing a fibonacci sequence
func Fibonacci(i int) []int {
	res := make([]int, i)

	if i > 0 {
		res[0] = 1

		if i > 1 {
			res[1] = 1
		}

		if i > 2 {
			for j := 2; j < i; j++ {
				res[j] = res[j-1] + res[j-2]
			}
		}
	}

	return res
}
