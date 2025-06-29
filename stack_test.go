package dsa

import (
	"testing"
)

func TestPush(t *testing.T) {
	var stack Stack[int32]

	stack.Push(9)
	stack.Push(10)

	if stack.Peek() != 10 {
		t.Errorf("After Push got: %v; Wanted: %v", stack.Peek(), 10)
	}

	if stack.Length != 2 {
		t.Errorf("Length incorrect; Got %v; Expected: %v", stack.Length, 2)
	}

}

func TestPop(t *testing.T) {

	var stack Stack[int32]

	stack.Push(4)
	stack.Push(5)
	stack.Push(15)
	stack.Pop()

	if stack.Peek() != 5 {
		t.Errorf("After pop got: %v; Wanted: %v", stack.Peek(), 5)
	}

	if stack.Length != 2 {
		t.Errorf("Length incorrect; Got %v; Expected: %v", stack.Length, 2)
	}

}

func TestPeek(t *testing.T) {
	var stack Stack[int32]

	if stack.Peek() != 0 {
		t.Errorf("Calling peek on an empty stack did not yield an zero value")
	}

	stack.Push(4)
	stack.Push(5)
	stack.Peek()
	if stack.Peek() != 5 {
		t.Errorf("Stack.Peek() did not return correct value")
	}
}

func TestPrint(t *testing.T) {
	var stack Stack[int32]

	stack.Push(4)
	stack.Push(5)
	stack.Push(6)

	if stack.Print() != "6 -> 5 -> 4\n" {
		t.Errorf("Print function did not print the correct string")
	}

}
