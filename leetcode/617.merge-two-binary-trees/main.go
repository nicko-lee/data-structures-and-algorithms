package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	/* traverse both trees at once in preorder fashion
	Note that at any point in time we only ever know about 2 nodes in both our trees, one from each tree
	*/
	// base conditions
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val // modifying t1. Alternatively, could also store result in t3
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

// This is a Depth First traversal
func printInOrder(head *TreeNode) {
	// check if currentNode is nil
	if head != nil {
		// traverse left subtree recursively
		printInOrder(head.Left)

		// print currentNode.Val
		fmt.Println(head.Val)

		// traverse right subtree recursively
		printInOrder(head.Right)
	}
}

func main() {
	// init tree 1 values and then pointers
	treeOneHead := TreeNode{Val: 1}
	treeOne3 := TreeNode{Val: 3}
	treeOne2 := TreeNode{Val: 2}
	treeOne5 := TreeNode{Val: 5}

	treeOneHead.Left = &treeOne3
	treeOneHead.Right = &treeOne2
	treeOne3.Left = &treeOne5

	// init tree 2 values and then pointers
	treeTwoHead := TreeNode{Val: 2}
	treeTwo1 := TreeNode{Val: 1}
	treeTwo3 := TreeNode{Val: 3}
	treeTwo4 := TreeNode{Val: 4}
	treeTwo7 := TreeNode{Val: 7}

	treeTwoHead.Left = &treeTwo1
	treeTwoHead.Right = &treeTwo3
	treeTwo1.Right = &treeTwo4
	treeTwo3.Right = &treeTwo7

	// printInOrder(&treeOneHead)
	// printInOrder(&treeTwoHead)
	printInOrder((mergeTrees(&treeOneHead, &treeTwoHead)))
}

/*
Note the hint tells us we should use pre-order traversal
The hint (The merging process must start from the root nodes of both trees)

Very useful to actually step thru the algo by hand as it is not immediately intuitive or easy to understand.

*/
