package minmax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	arr := []int{3, 1, 5, 4}

	res := Min(arr)

	assert.Equal(t, 1, res)
}

func TestMax(t *testing.T) {
	arr := []int{3, 1, 5, 4}

	res := Max(arr)

	assert.Equal(t, 5, res)
}
