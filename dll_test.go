package dsa

import "testing"

func initDLL() DLL[int] {
	var dll DLL[int]

	numbers := []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100}

	for _, num := range numbers {
		dll.InsertEnd(num)
	}

	return dll

}

func TestDLLInsertFront(t *testing.T) {
	var dll DLL[int]

	dll.InsertFront(5)
	if dll.first.item != 5 && dll.last.item != 5 {
		t.Errorf("Insert did not insert first item correctly")
	}

	dll.InsertFront(10)
	if dll.first.item != 10 || dll.last.item != 5 {
		t.Errorf("Did not insert 2 items correctly")
	}

	dll.InsertFront(15)
	if dll.first.item != 15 || dll.last.item != 5 {
		t.Errorf("Did not insert 3 items correctly")
	}

}

func TestDLLInsertEnd(t *testing.T) {
	var dll DLL[int]

	dll.InsertEnd(5)
	if dll.first.item != 5 && dll.last.item != 5 {
		t.Errorf("Insert did not insert first item correctly")
	}

	dll.InsertEnd(10)
	if dll.first.item != 5 || dll.last.item != 10 {
		t.Errorf("Did not insert 2 items correctly")
	}

}

func TestDLLLength(t *testing.T) {
	dll := initDLL()

	if dll.Length() != 20 {
		t.Errorf("Incorrect length reported; Expected %v; Got %v", 20, dll.Length())
	}
}

func TestDLLGet(t *testing.T) {
	dll := initDLL()

	err, item := dll.Get(8)
	if err != nil {
		t.Errorf("Error getting index")
	}
	if item != 45 {
		t.Errorf("Get did not get correct index; Expected value %v; Got %v", 45, item)
	}

	err, _ = dll.Get(69)
	if err == nil {
		t.Errorf("No error value trying to get index outside of range")
	}

	dll = DLL[int]{}
	err, _ = dll.Get(0)
	if err.Error() != "list is empty" {
		t.Errorf("Did not get proper error on empty list")
	}

}

func TestDLLRemove(t *testing.T) {
	dll := initDLL()

	err := dll.Remove(0)
	if err != nil || dll.first.item != 10 {
		t.Errorf("Removing an item did not remove the correct item")
	}
	if dll.Length() != 19 {
		t.Errorf("length did not update properly")
	}

	err = dll.Remove(9000)
	if err.Error() != "index out of range" {
		t.Errorf("did not error on index out of range")
	}

	err = dll.Remove(5)
	if err != nil {
		t.Errorf("failed to remove item from the list")
	}

	dll = DLL[int]{}

	err = dll.Remove(0)
	if err.Error() != "list is empty" {
		t.Errorf("did not give proper error for empty list")
	}

}
