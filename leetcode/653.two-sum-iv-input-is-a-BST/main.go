package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InsertNodeToTree(tree *TreeNode, node *TreeNode) {
	if tree == nil {
		panic("cannot insert into nil root")
	}

	if node.Val > tree.Val {
		if tree.Right == nil {
			tree.Right = node
		} else {
			InsertNodeToTree(tree.Right, node)
		}
	}
	if node.Val < tree.Val {
		if tree.Left == nil {
			tree.Left = node
		} else {
			InsertNodeToTree(tree.Left, node)
		}
	}
}

// note initTree is a variadic function
// it takes in a bunch of ints and returns a pointer to the root node
func InitTree(Vals ...int) (root *TreeNode) {
	// start with creating the root node and inserting the first value of the variadic array into it
	rootNode := TreeNode{Val: Vals[0]}
	for _, Val := range Vals {
		node := TreeNode{Val: Val}
		InsertNodeToTree(&rootNode, &node)
	}
	return &rootNode
}

func searchTree(val int, root *TreeNode) bool {

	// start at root node and traverse downwards first (DFS)
	// check left child - init, condition, then go to next node
	for currentNode := root; root.Left != nil; currentNode = currentNode.Left {
		fmt.Println("I am in the left for loop")
		fmt.Println(currentNode.Val)
		// check if exists in children both Left and Right
		if currentNode.Left.Val == val || currentNode.Right.Val == val {
			return true
		}
		if currentNode.Left == nil {
			break
		}
	}

	// check right child - init, condition, then go to next node
	for currentNode := root; root.Right != nil; currentNode = currentNode.Right {
		fmt.Println("I am in the right for loop")
		fmt.Println(currentNode.Val)
		// check if exists in children both Left and Right
		if currentNode.Left.Val == val || currentNode.Right.Val == val {
			return true
		}
		if currentNode.Right == nil {
			break
		}
	}

	// fmt.Println(val)

	return false
}

func findTarget(root *TreeNode, target int) int {
	// start at root and deduct target from root
	diff := target - root.Val

	// find the other number in the tree

	// if it doesn't exist repeat on the next consecutive node for the remainder

	fmt.Println(diff)
	return diff
}

func main() {
	treeNode := InitTree(5, 3, 6, 2, 4, 7)

	// target := 9
	// fmt.Println(treeNode)
	// findTarget(treeNode, target)
	searchTree(5, treeNode)
}
