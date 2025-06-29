package dsa

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// Insert adds a new value into the binary search tree (BST) while maintaining its structural properties.
func (tree *BST) Insert(value int) *BST {

	// TODO Implement logic for it to insert in optimal location rather than just at the ends every time

	currentNode := tree

	// When the tree is initialized the first value is 0. When we add a node we want it to become the root
	if value != 0 && tree.Left == nil && tree.Right == nil && tree.Value == 0 {
		tree.Value = value
		return tree
	}

	for {
		if value < currentNode.Value && currentNode.Left != nil {
			currentNode = currentNode.Left
			continue
		} else if (value > currentNode.Value || currentNode.Value == value) && currentNode.Right != nil {
			currentNode = currentNode.Right
			continue
		}

		if value < currentNode.Value && currentNode.Left == nil {
			currentNode.Left = &BST{value, nil, nil}
			break
		} else if (value > currentNode.Value || value == currentNode.Value) && currentNode.Right == nil {
			currentNode.Right = &BST{value, nil, nil}
			break
		}
	}

	return tree
}

// Contains checks if the specified value exists in the binary search tree (BST) and returns true if found, otherwise false.
func (tree *BST) Contains(value int) bool {

	currentNode := tree

	for {

		if currentNode.Value == value {
			return true
		}

		if value < currentNode.Value {
			if currentNode.Left == nil {
				return false
			} else {
				currentNode = currentNode.Left
			}
		} else if value > currentNode.Value {
			if currentNode.Right == nil {
				return false
			} else {
				currentNode = currentNode.Right
			}
		}
	}
}

// Remove deletes a node with the given value from the BST and adjusts the tree structure accordingly.
func (tree *BST) Remove(value int) *BST {

	if !(tree.Contains(value)) {
		return nil
	}

	currentNode := tree
	var prevNode *BST = nil

	for {

		// When we find the correct node
		if currentNode.Value == value {

			// For root node
			if prevNode == nil {
				//Deleting root node if only node
				if currentNode.Left == nil && currentNode.Right == nil {
					return nil
				}

				if currentNode.Right != nil {

					rightSideMin := currentNode.Right.GetMin()
					if currentNode.Right.Value == rightSideMin {
						currentNode.Right = nil
					} else {
						currentNode.Remove(rightSideMin)
					}
					currentNode.Value = rightSideMin
					return tree
				}

				if currentNode.Left != nil && currentNode.Right == nil {
					*tree = *currentNode.Left
					return tree
				}
			}

			// If the node to delete has no children nodes
			if currentNode.Left == nil && currentNode.Right == nil {

				if currentNode.Value < prevNode.Value {
					prevNode.Left = nil
				} else if currentNode.Value > prevNode.Value {
					prevNode.Right = nil
				}
				return tree
			}

			// If the node to delete has one child node
			if (currentNode.Left != nil && currentNode.Right == nil) || (currentNode.Left == nil && currentNode.Right != nil) {
				if currentNode.Left != nil {
					if currentNode.Value < prevNode.Value {
						prevNode.Left = currentNode.Left
					} else if currentNode.Value > prevNode.Value {
						prevNode.Right = currentNode.Left
					}
				} else if currentNode.Right != nil {
					if currentNode.Value < prevNode.Value {
						prevNode.Left = currentNode.Right
					} else if currentNode.Value > prevNode.Value {
						prevNode.Right = currentNode.Right
					}
				}
				return tree
			}

			// If the node to delete has multiple children
			if currentNode.Left != nil && currentNode.Right != nil {
				rightSideMin := currentNode.Right.GetMin()
				if currentNode.Right.Value == rightSideMin {
					currentNode.Right = nil
				} else {
					currentNode.Remove(rightSideMin)
				}
				currentNode.Value = rightSideMin

				return tree
			}

		}

		if currentNode.Left == nil && currentNode.Right == nil {
			return tree
		}

		prevNode = currentNode

		// For navigation
		if value < currentNode.Value && currentNode.Left != nil {
			currentNode = currentNode.Left
		} else if value > currentNode.Value && currentNode.Right != nil {
			currentNode = currentNode.Right
		}
	}
}

// GetMin traverses the leftmost branch of the binary search tree (BST) and returns the minimum value found in the tree.
func (tree *BST) GetMin() int {

	currentNode := tree

	for {
		if currentNode.Left == nil {
			break
		}
		currentNode = currentNode.Left
	}
	return currentNode.Value
}

// GetMax traverses the rightmost branch of the binary search tree (BST) and returns the maximum value in the tree.
func (tree *BST) GetMax() int {

	currentNode := tree

	for {
		if currentNode.Right == nil {
			break
		}
		currentNode = currentNode.Right
	}
	return currentNode.Value
}

func addValues(tree *BST, values *[]int) {

	*values = append(*values, tree.Value)

	if tree.Left != nil {
		addValues(tree.Left, values)
	}

	if tree.Right != nil {
		addValues(tree.Right, values)
	}
}

func getSliceMedian(slice []int) int {
	return len(slice) / 2
}

func (tree *BST) Balance() {

	// TODO Implement test

	var values []int
	addValues(tree, &values)

	root := getSliceMedian(values)
	tree.Value = values[root]
	tree.Left = nil
	tree.Right = nil

	for i, value := range values {
		if i == root {
			continue
		}
		tree.Insert(value)
	}

	values = []int{}
	addValues(tree, &values)
}
