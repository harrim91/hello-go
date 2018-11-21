package arrmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	addOne := func(n int) int {
		return n + 1
	}
	arr := []int{1, 2, 3}

	expected := []int{2, 3, 4}

	res := Map(addOne, arr)

	assert.Equal(t, expected, res)
}
