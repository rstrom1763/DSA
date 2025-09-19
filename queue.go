package dsa

type Queue[T any] struct {
	sll *SLL[T]
}

func (q *Queue[T]) init() {

	q.sll = &SLL[T]{
		length: 0,
		first:  nil,
		last:   nil,
	}
	q.sll.length = 0
	q.sll.first = nil
	q.sll.last = nil
}

// Queue adds an item to the end of the queue and increments the queue's length by 1.
func (q *Queue[T]) Queue(item T) {
	if q.sll == nil {
		q.init()
	}
	q.sll.InsertEnd(item)
}

// Dequeue removes and returns the first item from the queue. Updates the queue's first node and decrements its length.
func (q *Queue[T]) Dequeue() {
	_ = q.sll.Remove(0)
}

// Size returns the number of elements currently in the queue.
func (q *Queue[T]) Size() int64 {
	if q.sll == nil {
		q.init()
	}
	return q.sll.Length()
}

// Peek returns the first item of the queue without removing it. If the queue is empty, it returns the zero value of the type.
func (q *Queue[T]) Peek() T {
	if q.sll == nil {
		q.init()
	}
	if q.Size() == 0 {
		var zero T
		return zero
	}
	return q.sll.getFirst()
}
