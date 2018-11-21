package average

// Average returns the average of an array of float64 numbers
func Average(numbers []float64) (avg float64) {
	sum := 0.0
	switch len(numbers) {
	case 0:
		avg = 0.0
	default:
		for _, v := range numbers {
			sum += v
		}
		avg = sum / float64(len(numbers))
	}

	return
}
