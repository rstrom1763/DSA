package dsa

import "fmt"

type DLL[T any] struct {
	first  *dllNode
	last   *dllNode
	length int64
}

type dllNode struct {
	item any
	next *dllNode
	prev *dllNode
}

// InsertFront Insert at the front of the dll
func (dll *DLL[T]) InsertFront(item T) {

	defer func() {
		dll.length += 1
	}()

	if dll.Length() == 0 {
		newNode := &dllNode{item: item, prev: nil, next: nil}
		dll.first = newNode
		dll.last = newNode
		return
	}

	if dll.Length() == 1 {
		dll.first = &dllNode{item: item, prev: nil, next: dll.first}
		dll.first.next.prev = dll.first
		return
	}

	if dll.Length() > 1 {
		dll.first = &dllNode{item: item, prev: nil, next: dll.first.next}
		dll.first.next.prev = dll.first
		return
	}

}

func (dll *DLL[T]) InsertEnd(item T) {

	defer func() {
		dll.length += 1
	}()

	newNode := &dllNode{item: item, prev: dll.last, next: nil}

	if dll.Length() == 0 {
		dll.first = newNode
		dll.last = newNode
		return
	}

	if dll.Length() > 0 {
		newNode.prev = dll.last
		dll.last.next = newNode
		dll.last = newNode
	}
}

func (dll *DLL[T]) Length() int64 {
	return dll.length
}

func (dll *DLL[T]) Get(index int64) (error, T) {
	
	if index <= (dll.Length() / 2) {
		var i int64 = 0
		var zero T

		if dll.first == nil {
			return fmt.Errorf("list is empty"), zero
		}

		currentNode := dll.first

		if index >= dll.Length() {
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
	} else {
		var i int64 = dll.Length() - 1
		var zero T

		if dll.first == nil {
			return fmt.Errorf("list is empty"), zero
		}

		currentNode := dll.last

		if index >= dll.Length() {
			return fmt.Errorf("index out of range"), zero
		}

		for {
			if index == i {
				return nil, currentNode.item.(T)
			} else {
				currentNode = currentNode.prev
				i -= 1
			}
		}
	}

}

func (dll *DLL[T]) Remove(index int64) error {

	var i int64 = 0

	if dll.first == nil {
		return fmt.Errorf("list is empty")
	}

	currentNode := dll.first

	if index >= dll.Length() {
		return fmt.Errorf("index out of range")
	}

	for {

		if index == 0 && dll.Length() > 1 {
			dll.first = dll.first.next
			dll.length -= 1
			return nil
		}

		if dll.Length() == 1 {
			dll.first = nil
			dll.last = nil
			dll.length -= 1
			return nil
		}

		if index == i {
			currentNode.prev.next = currentNode.next
			currentNode.next.prev = currentNode.prev
			currentNode.next = nil
			dll.length -= 1
			return nil
		} else {
			currentNode = currentNode.next
			i += 1
		}
	}
}
