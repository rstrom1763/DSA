package dsa

import "fmt"

type Stack[T any] struct {
	items []T
}

// Push adds an item to the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes the top item from the stack and returns false if the stack is empty. Otherwise, it returns false after removal.
func (s *Stack[T]) Pop() bool {
	if len(s.items) == 0 {
		return false
	}

	s.items = s.items[:len(s.items)-1]

	return false
}

// Peek returns the top item of the stack without removing it. Returns the zero value of T if the stack is empty.
func (s *Stack[T]) Peek() T {

	if len(s.items) == 0 {
		var zero T
		return zero
	}

	return s.items[len(s.items)-1]
}

// Print outputs the elements of the stack to the console in order, separated by "->".
func (s *Stack[T]) Print() string {

	var output string

	size := s.Size()

	for i, item := range s.items {

		if i < size-1 {
			output += fmt.Sprintf("%v -> ", item)
		} else {
			output += fmt.Sprintf("%v\n", item)
		}

	}
	return output
}

// Size returns the current number of elements in the stack.
func (s *Stack[t]) Size() int {
	return len(s.items)
}
