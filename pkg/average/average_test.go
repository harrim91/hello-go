package average

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnsZero(t *testing.T) {
	numbers := [0]float64{}
	assert.Equal(t, 0.0, Average(numbers[:]))
}

func TestReturnsAverage(t *testing.T) {
	numbers := [2]float64{5.0, 0.0}
	assert.Equal(t, 2.5, Average(numbers[:]))
}
