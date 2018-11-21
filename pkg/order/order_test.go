package order

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnsUnorderedParamsInOrder(t *testing.T) {
	first, second := Order(7, 2)
	assert.Equal(t, first, 2)
	assert.Equal(t, second, 7)
}

func TestReturnsOrderedParamsInOrder(t *testing.T) {
	first, second := Order(2, 7)
	assert.Equal(t, first, 2)
	assert.Equal(t, second, 7)
}
