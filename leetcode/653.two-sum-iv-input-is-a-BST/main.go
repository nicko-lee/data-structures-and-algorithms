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

func stringify(tree *TreeNode, level int) {
	if tree != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(tree.Right, level)
		fmt.Printf(format+"%d\n", tree.Val)
		stringify(tree.Left, level)
	}
}

func main() {
	treeNode := InitTree(5, 3, 6, 2, 4, 7)
	treeNode := InitTree(5, 4, 2, 6, 3, 7)
	stringify(treeNode, 5)

	// target := 9
	// fmt.Println(treeNode)
	// findTarget(treeNode, target)
	// searchTree(5, treeNode)
}
