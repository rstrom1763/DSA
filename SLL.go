package dsa

import "fmt"

type SLL[T any] struct {
	first  *sllNode
	last   *sllNode
	length int64
}

type sllNode struct {
	item any
	next *sllNode
}

// InsertFront Insert at the front of the SLL
func (sll *SLL[T]) InsertFront(item T) {

	defer func() {
		sll.length += 1
	}()

	if sll.first == nil {
		newNode := &sllNode{item: item, next: nil}
		sll.first = newNode
		sll.last = newNode
		return
	}

	if sll.first != nil {
		sll.first = &sllNode{item: item, next: sll.first.next}
		return
	}

}

func (sll *SLL[T]) InsertEnd(item T) {

	defer func() {
		sll.length += 1
	}()

	newNode := &sllNode{item: item, next: nil}

	if sll.first == nil {
		sll.first = newNode
		sll.last = newNode
		return
	}

	sll.last.next = newNode
	sll.last = newNode

}

func (sll *SLL[T]) Length() int64 {
	return sll.length
}

func (sll *SLL[T]) Get(index int64) (error, T) {
	var i int64 = 0
	var zero T

	if sll.first == nil {
		return fmt.Errorf("list is empty"), zero
	}

	currentNode := sll.first

	if index >= sll.Length() {
		return fmt.Errorf("index out of range"), zero
	}

	for {
		if index == i {
			return nil, currentNode.item.(T)
		} else {
			currentNode = currentNode.next
			i += 1
		}
	}
}

func (sll *SLL[T]) Remove(index int64) error {
	var i int64 = 0

	if sll.first == nil {
		return fmt.Errorf("list is empty")
	}

	currentNode := sll.first
	var prevNode *sllNode

	if index >= sll.Length() {
		return fmt.Errorf("index out of range")
	}

	for {

		if index == 0 {
			sll.first = sll.first.next
			sll.length -= 1
			return nil
		}

		if index == i {
			prevNode.next = currentNode.next
			currentNode.next = nil
			sll.length -= 1
			return nil
		} else {
			prevNode = currentNode
			currentNode = currentNode.next
			i += 1
		}
	}
}
