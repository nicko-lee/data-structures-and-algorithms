package main

import "fmt"

// Do not edit the class below except for
// the insert, contains, and remove methods.
// Feel free to add new properties and methods
// to the class.
type BST struct {
	value int

	left  *BST
	right *BST
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

// deletion/removal
// the most difficult method relatively speaking - cos quite a few edge cases
// at high level think of removal as 2 step process. First u must find the node. Then actually remove it
// func (tree *BST) Remove(value int) *BST {
// 	// Write your code here.
// 	currentNode := tree
// 	parentNode := currentNode

// 	for currentNode != nil {
// 		if value < currentNode.value {
// 			parentNode = currentNode
// 			currentNode = currentNode.left
// 		} else if value > currentNode.value {
// 			parentNode = currentNode
// 			currentNode = currentNode.right
// 		} else { // this is the bread and butter bit. This is the case where we found it and going to remove it
// 			// note within this case there are a few futher subcases
// 			// case 1: dealing with a node with 2 children nodes (the tricky case)
// 			// a simple rule to deal with this case is to grab the smallest value (leftmost) in right subtree to take its place
// 			// this will always result in a valid BST
// 			if currentNode.left != nil && currentNode.right != nil {
// 				// grabbing the smallest value of the right subtree and placing it in its place
// 				// currentNode.value = currentNode.right.getMinValue()
// 				// currentNode.right.Remove(currentNode.value)
// 			}

// 			// case 2: node has 0 children (easy case)

// 			// case 3: node has 1 child (easy case just promote the child)

// 		}
// 	}
// 	// Do not edit the return statement of this method.
// 	return tree
// }

// internal recursive function to print a tree
func stringify(tree *BST, level int) {
	if tree != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(tree.left, level)
		fmt.Printf(format+"%d\n", tree.value)
		stringify(tree.right, level)
	}
}

func main() {
	tree := &BST{value: 6}
	fmt.Println(tree)
	fmt.Println(tree.Insert(5).Insert(1).Insert(7).Insert(0).Insert(3).Insert(8))
	stringify(tree, 4)
}

/* -------- NOTES -----------

High level algorithm steps

For deletion/removal

1. search
	- while searching keep track of current node and previous node (i.e the parent)
	- search left
	- search right
	- found - now look at step #2: delete
2. delete
	- case 1 = 2 children
		- use promotion but must search for smallest (leftmost) value of right subtree
		- when found put that value in current slot to take it's place
		- clean up and remove that value u just plucked from the bottom leaf
	- case 2 = 1 child (easy) - use direct promotion from 1 child below
	- case 3 = 0 children (easy) - easy just remove and that's it
	- note: for deleting - there 2 things to delete
		- delete the node itself
		- delete the pointer in the parent
		- ref: https://www.youtube.com/watch?v=gcULXE7ViZw
		- in particular at this timestamp: https://youtu.be/gcULXE7ViZw?t=586


Other good resources
1. https://appliedgo.net/bintree/

*/
