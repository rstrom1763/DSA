package dsa

import "testing"

func initSLL() SLL[int] {
	var sll SLL[int]

	numbers := []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100}

	for _, num := range numbers {
		sll.InsertEnd(num)
	}

	return sll

}

func TestInsertFront(t *testing.T) {
	var sll SLL[int]

	sll.InsertFront(5)
	if sll.first.item != 5 && sll.last.item != 5 {
		t.Errorf("Insert did not insert first item correctly")
	}

	sll.InsertFront(10)
	if sll.first.item != 10 || sll.last.item != 5 {
		t.Errorf("Did not insert 2 items correctly")
	}

}

func TestInsertEnd(t *testing.T) {
	var sll SLL[int]

	sll.InsertEnd(5)
	if sll.first.item != 5 && sll.last.item != 5 {
		t.Errorf("Insert did not insert first item correctly")
	}

	sll.InsertEnd(10)
	if sll.first.item != 5 || sll.last.item != 10 {
		t.Errorf("Did not insert 2 items correctly")
	}

}

func TestLength(t *testing.T) {
	sll := initSLL()

	if sll.Length() != 20 {
		t.Errorf("Incorrect length reported; Expected %v; Got %v", 20, sll.Length())
	}
}

func TestGet(t *testing.T) {
	sll := initSLL()

	err, item := sll.Get(8)
	if err != nil {
		t.Errorf("Error getting index")
	}
	if item != 45 {
		t.Errorf("Get did not get correct index; Expected value %v; Got %v", 45, item)
	}

	err, _ = sll.Get(69)
	if err == nil {
		t.Errorf("No error value trying to get index outside of range")
	}

	sll = SLL[int]{}
	err, _ = sll.Get(0)
	if err.Error() != "list is empty" {
		t.Errorf("Did not get proper error on empty list")
	}

}

func TestSLLRemove(t *testing.T) {
	sll := initSLL()

	err := sll.Remove(0)
	if err != nil || sll.first.item != 10 {
		t.Errorf("Removing an item did not remove the correct item")
	}
	if sll.Length() != 19 {
		t.Errorf("length did not update properly")
	}

	err = sll.Remove(9000)
	if err.Error() != "index out of range" {
		t.Errorf("did not error on index out of range")
	}

	err = sll.Remove(5)
	if err != nil {
		t.Errorf("failed to remove item from the list")
	}

	sll = SLL[int]{}

	err = sll.Remove(0)
	if err.Error() != "list is empty" {
		t.Errorf("did not give proper error for empty list")
	}

}
