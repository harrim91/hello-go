package minmax

// Min returns the smallest value in an array
func Min(arr []int) (res int) {
	for i, n := range arr {
		if i == 0 || n < res {
			res = n
		}
	}
	return
}

// Max returns th largest value in an array
func Max(arr []int) (res int) {
	for _, n := range arr {
		if n > res {
			res = n
		}
	}
	return
}
