package dsa

import "fmt"

type Queue[T any] struct {
	First  *Node
	Last   *Node
	Length int32
}

type Node struct {
	Item any
	Next *Node
}

// Queue adds an item to the end of the queue and increments the queue's length by 1.
func (q *Queue[T]) Queue(item T) {

	// Increment length by 1
	defer func() { q.Length += 1 }()

	node := Node{
		Item: item,
		Next: nil,
	}

	// If first item in the queue
	if q.Length == 0 {
		q.First = &node
		q.Last = &node
		return
	}

	// Set the next field of the last to be the new item as it is now the new last
	q.Last.Next = &node
	q.Last = &node // Set last to the new node

}

// Dequeue removes and returns the first item from the queue. Updates the queue's first node and decrements its length.
func (q *Queue[T]) Dequeue() any {

	// Decrement length by 1
	defer func() { q.Length -= 1 }()

	// Capture the value of the item to be dequeued
	returnValue := q.First.Item

	// Set the next item in the queue to be first
	q.First = q.First.Next

	return returnValue

}

// Size returns the number of elements currently in the queue.
func (q *Queue[T]) Size() int32 {
	return q.Length
}

// Print outputs all the items in the queue in their current order, separated by '->', ending with a newline.
func (q *Queue[T]) Print() string {
	item := q.First
	var output string
	for {
		if item.Next != nil {
			output += fmt.Sprintf("%v -> ", item.Item)
			item = item.Next
		} else { // For the last item in the queue
			output += fmt.Sprintf("%v\n", item.Item)
			return output
		}
	}
}

// Peek returns the first item of the queue without removing it. If the queue is empty, it returns the zero value of the type.
func (q *Queue[T]) Peek() any {

	if q.First == nil {
		return nil
	}

	return q.First.Item
}

// PeekLast returns the last item in the queue without removing it. If the queue is empty, it returns nil.
func (q *Queue[T]) PeekLast() any {
	if q.Size() > 0 {
		return q.Last.Item
	} else {
		return nil
	}
}
