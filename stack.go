package dsa

import "fmt"

type Stack[T any] struct {
	First  *node
	Length int32
}

type node struct {
	Item any
	Next *node
}

// Push adds an item to the stack
func (s *Stack[T]) Push(item T) {
	newNode := node{Item: item, Next: s.First}
	s.First = &newNode
	s.Length += 1
}

// Pop removes the top item from the stack and returns false if the stack is empty. Otherwise, it returns false after removal.
func (s *Stack[T]) Pop() {

	if s.Length > 1 {
		s.First = s.First.Next
		s.Length -= 1
		return
	}

	if s.Length == 1 {
		s.First = nil
		s.Length -= 1
		return
	}

	if s.Length == 0 {
		return
	}

}

// Peek returns the top item of the stack without removing it. Returns the zero value of T if the stack is empty.
func (s *Stack[T]) Peek() T {

	var zero T

	if s.Length > 0 {
		return s.First.Item.(T)
	}

	if s.Length == 0 {
		return zero
	}

	return zero

}

// Print outputs the elements of the stack to the console in order, separated by "->".
func (s *Stack[T]) Print() string {

	var output string
	currentItem := s.First

	for {
		if currentItem.Next == nil {
			output += fmt.Sprintf("%v\n", currentItem.Item)
			return output
		} else {
			output += fmt.Sprintf("%v -> ", currentItem.Item)
		}
		currentItem = currentItem.Next
	}
}

// Size returns the current number of elements in the stack.
func (s *Stack[t]) Size() int32 {
	return s.Length
}
