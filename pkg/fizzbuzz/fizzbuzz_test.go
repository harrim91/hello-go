package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnsStringifiedIfNormalNumber(t *testing.T) {
	res := FizzBuzz(2)
	assert.Equal(t, "2", res)
}

func TestReturnsFizzIfMultipleOf3(t *testing.T) {
	res := FizzBuzz(3)
	assert.Equal(t, "Fizz", res)
}

func TestReturnsBuzzIfMultipleOf5(t *testing.T) {
	res := FizzBuzz(5)
	assert.Equal(t, "Buzz", res)
}

func TestReturnsFizzBuzzIfMultipleOf15(t *testing.T) {
	res := FizzBuzz(15)
	assert.Equal(t, "FizzBuzz", res)
}
