package dsa

import (
	"testing"
)

func initTree() *BST {
	var tree BST

	var values = []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100}
	for _, num := range values {
		tree.Insert(num)
	}
	tree.Balance()
	return &tree
}

func TestContains(t *testing.T) {

	tree := initTree()
	tree.Balance()

	//fmt.Printf("Current root node: %v\n", tree.Value)

	if tree.Contains(10) == false {
		t.Errorf("Tree contains failed")
	}

}

func TestInsert(t *testing.T) {
	var tree BST

	tree.Insert(5)
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(1)

	if tree.Contains(10) != true {
		t.Errorf("Tree insert failed")
	}

}

func TestRemove(t *testing.T) {

	tree := initTree()

	tree.Remove(15)
	if tree.Contains(15) == true {
		t.Errorf("BST remove failed; Failed to remove a node")
	}

	tree.Remove(55)
	if tree.Contains(55) == true {
		t.Errorf("BST remove failed; Failed to delete root node")
	}

	tree.Remove(100)
	if tree.Contains(100) == true {
		t.Errorf("BST remove failed; Failed to remove a node")
	}

	tree.Remove(60)
	if tree.Contains(60) == true {
		t.Errorf("BST remove failed; Failed to remove a node")
	}

	// TODO tests for specifically deleting nodes with 0, 1 and 2 child nodes

}
 
func TestGetMin(t *testing.T) {
	var tree BST

	tree.Insert(5)
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(1)

	tree.Remove(0)

	expected := 1
	value := tree.GetMin()

	if value != expected {
		t.Errorf("BST GetMin(); Got: %v; Expected: %v", value, expected)
	}

}

func TestGetMax(t *testing.T) {
	var tree BST

	tree.Insert(5)
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(1)

	expected := 15
	value := tree.GetMax()

	if value != expected {
		t.Errorf("BST GetMax(); Got: %v; Expected: %v", value, expected)
	}

}

func TestBalance(t *testing.T) {

	tree := initTree()

	if tree.Value != 55 {
		t.Errorf("Did not correctly balance tree;Expected 55;Got %v\n", tree.Value)
	}

}
