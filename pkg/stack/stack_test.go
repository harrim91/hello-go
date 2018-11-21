package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStackIsEmpty(t *testing.T) {
	s := new(Stack)
	expected := [10]int{}
	assert.Equal(t, expected, s.Data)
}

func TestPushAddsIntToStack(t *testing.T) {
	s := new(Stack)
	s.Push(1)
	expected := [10]int{1}
	assert.Equal(t, expected, s.Data)

	s.Push(2)
	expected = [10]int{1, 2}
	assert.Equal(t, expected, s.Data)
}

func TestPopReturnsLastIntFromStack(t *testing.T) {
	s := new(Stack)
	s.Push(1)
	s.Push(2)

	res := s.Pop()
	assert.Equal(t, 2, res)

	res = s.Pop()
	assert.Equal(t, 1, res)
}

func TestStringifyReturnsStringOfStackData(t *testing.T) {
	s := new(Stack)
	s.Push(1)
	s.Push(2)

	res := s.Stringify()
	assert.Equal(t, "[0:1, 1:2]", res)

	fmt.Println(s.Stringify())
}
