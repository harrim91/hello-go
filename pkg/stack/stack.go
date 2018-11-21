package stack

import (
	"strconv"
)

// Stack manages a stack of numbers
type Stack struct {
	Data [10]int
	i    int
}

// Push adds a number to the stack
func (s *Stack) Push(i int) {
	s.Data[s.i] = i
	s.i++
}

// Pop removes and returns the last number from the stack
func (s *Stack) Pop() int {
	s.i--
	return s.Data[s.i]
}

// Stringify returns a string representation of the stack
func (s Stack) Stringify() string {
	if s.i == 0 {
		return ""
	}

	str := "["

	for i := 0; i < s.i; i++ {
		str += strconv.Itoa(i) + ":" + strconv.Itoa(s.Data[i])
		if i != s.i-1 {
			str += ", "
		}
	}

	str += "]"

	return str
}
