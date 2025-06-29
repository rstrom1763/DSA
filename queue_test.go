package dsa

import (
	"testing"
)

func TestQueue(t *testing.T) {
	var test Queue[int32]

	test.Queue(5)
	test.Queue(10)
	test.Queue(15)

	var expected int32 = 5

	if test.Peek() != expected {
		t.Errorf("Queueing a value; Expected: %v; Got: %v", expected, test.Peek())
	}

}

func TestSize(t *testing.T) {
	var test Queue[int32]

	test.Queue(5)
	test.Queue(10)
	test.Queue(15)

	if test.Size() != 3 {
		t.Errorf("Expected size: %v; Got: %v", 3, test.Size())
	}

}

func TestQueuePeek(t *testing.T) {
	var test Queue[int32]

	if test.Peek() != nil {
		t.Errorf("Peeking a value; Expected: nil; Got: %v", test.Peek())
	}

	test.Queue(5)
	test.Queue(10)
	test.Queue(15)

	var expected int32 = 5

	if test.Peek() != expected {
		t.Errorf("Peeking a value; Expected: %v; Got: %v", expected, test.Peek())
	}
}

func TestPeekLast(t *testing.T) {
	var test Queue[int32]

	if test.PeekLast() != nil {
		t.Errorf("Peeking last with empty queue; Expected: %v; Got: %v", "nil", test.PeekLast())
	}

	test.Queue(5)
	test.Queue(10)
	test.Queue(15)

	var expected int32 = 15

	if test.PeekLast() != expected {
		t.Errorf("Peeking last; Expected: %v; Got: %v", expected, test.PeekLast())
	}

}

func TestDequeue(t *testing.T) {
	var test Queue[int32]

	test.Queue(5)
	test.Queue(10)
	test.Queue(15)

	var expected int32 = 10

	test.Dequeue()

	if test.Peek() != expected {
		t.Errorf("Peeking; Expected: %v; Got: %v", expected, test.Peek())
	}
}

func TestQueuePrint(t *testing.T) {
	var test Queue[int32]

	test.Queue(5)
	test.Queue(10)
	test.Queue(15)

	expected := "5 -> 10 -> 15\n"

	if test.Print() != expected {
		t.Errorf("Printing; Expected: %v; Got: %v", expected, test.Print())
	}

}
