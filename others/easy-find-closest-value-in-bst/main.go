package main

import "fmt"

type BST struct {
	value int

	left  *BST
	right *BST
}

// Assumes there can only be one closest value. Doesn't take into account negative values of target - i
func (tree *BST) FindClosestValue(target int) int {
	// Write your code here.
	// Try and look for target in tree
	if tree.Contains(target) {
		return target
	} else {
		// Increment by 1 each time till find something
		// check (i + 1) and (i - 1) cases
		i := 1
		for {
			if tree.Contains(target + i) {
				return target + i
			}
			if tree.Contains(target - i) {
				return target - i
			}
			i++
		}
	}
}

func NewBST(root int, values ...int) *BST {
	tree := &BST{value: root}
	for _, value := range values {
		tree.Insert(value)
	}
	return tree
}

// Average: O(log(n)) time | O(1) space
// Worst: O(n) time | O(1) space
// insertion
func (tree *BST) Insert(value int) *BST {
	// Write your code here.
	// note: we call insertion method on the root node

	// 1. keep track of currentNode. Currently we are at root node
	currentNode := tree

	// 2. loop forever until break out
	for true {
		// explore left
		if value < currentNode.value {
			// check if we reached the end of the tree
			if currentNode.left == nil {
				// insert a value in open empty slot
				currentNode.left = &BST{value: value}
				break
			} else {
				// means we still have nodes to explore
				currentNode = currentNode.left // traversing the next node
			}
		} else if value >= currentNode.value {
			// explore right
			if currentNode.right == nil {
				currentNode.right = &BST{value: value}
				break
			} else {
				currentNode = currentNode.right
			}
		}
	}

	// Do not edit the return statement of this method
	return tree
}

// internal recursive function to print a tree
// source: https://flaviocopes.com/golang-data-structure-binary-search-tree/
func stringify(tree *BST, level int) {
	if tree != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(tree.right, level)
		fmt.Printf(format+"%d\n", tree.value)
		stringify(tree.left, level)
	}
}

// Average: O(log(n)) time | O(1) space
// Worst: O(n) time | O(1) space
// searching
// very simple similar to insert but simpler
func (tree *BST) Contains(value int) bool {
	// Write your code here.

	currentNode := tree

	// this is a while loop whereby u loop until u reach the end of a branch
	for currentNode != nil {
		// check if go left - if meets this check means the value ur looking for is in left subtree
		// can eliminate right completely
		// note in this 3 branch if else block. It is OR. Means only one of the code blocks will ever run
		// after all the value can only either be greater than or less than or equal too. Not ever more than one
		if value < currentNode.value {
			currentNode = currentNode.left
		} else if value > currentNode.value { // check if go right
			currentNode = currentNode.right
		} else { // if it is not less or greater means u have found it
			return true
		}
	}

	// Do not edit the return statement of this method
	return false
}

func main() {
	tree := &BST{value: 10}
	fmt.Println(tree.Insert(5).Insert(15).Insert(2).Insert(5).Insert(13).Insert(22).Insert(1).Insert(14))
	fmt.Println(NewBST(5, 15, 2, 5, 13, 22, 1, 14))
	stringify(tree, 7)
	fmt.Println(tree.Contains(5))
	fmt.Println(tree.FindClosestValue(12))
}

/* -------- NOTES -----------

Alright at the start I didn't even know how to get the input in and tried to "instantiate" 9 different BST structs
and tried to get them to point to each other.

But then I went off to do the BST construction question to get a much better way of doing this. Simply by adding an Insert
method for the type BST.

*/
