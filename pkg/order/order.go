package order

// Order returns its two arguments in numerical ascending order
func Order(x, y int) (int, int) {
	if x <= y {
		return x, y
	}
	return y, x
}
