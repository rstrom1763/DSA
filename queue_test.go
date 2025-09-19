package dsa

import (
	"testing"
)

// Creates a queue with a few items in it for testing purposes
func initTestQueue() Queue[int32] {
	var test Queue[int32]

	test.Queue(5)
	test.Queue(10)
	test.Queue(15)

	return test
}

func TestQueue(t *testing.T) {

	test := initTestQueue()

	var expected int32 = 5

	if test.Peek() != expected {
		t.Errorf("Queueing a value; Expected: %v; Got: %v", expected, test.Peek())
	}

}

func TestQueueSize(t *testing.T) {
	var test Queue[int32]

	if test.Size() != 0 {
		t.Errorf("Empty queue does not have size of 0")
	}

	test.Queue(5)
	test.Queue(10)
	test.Queue(15)

	if test.Size() != 3 {
		t.Errorf("Expected size: %v; Got: %v", 3, test.Size())
	}

}

func TestQueuePeek(t *testing.T) {
	var test Queue[int32]

	if test.Peek() != 0 && test.Size() != 0 {
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
