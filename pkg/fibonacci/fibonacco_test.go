package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	var res []int

	res = Fibonacci(0)
	assert.Equal(t, []int{}, res)

	res = Fibonacci(1)
	assert.Equal(t, []int{1}, res)

	res = Fibonacci(2)
	assert.Equal(t, []int{1, 1}, res)

	res = Fibonacci(5)
	assert.Equal(t, []int{1, 1, 2, 3, 5}, res)

	res = Fibonacci(10)
	assert.Equal(t, []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}, res)
}
