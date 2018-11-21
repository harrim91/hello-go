package fizzbuzz

import (
	strconv "strconv"
)

// FizzBuzz tells you if a function is divisible by 3, 5 or both
func FizzBuzz(n int64) string {
	if n%15 == 0 {
		return "FizzBuzz"
	}

	if n%3 == 0 {
		return "Fizz"
	}

	if n%5 == 0 {
		return "Buzz"
	}

	return strconv.FormatInt(n, 10)
}
