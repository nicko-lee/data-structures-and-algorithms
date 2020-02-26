package main

import (
	"fmt"
	"strconv"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	s := ""
	currentnode := head
	for currentnode.Next != nil {
		// build a string from the linked list
		s += strconv.Itoa(currentnode.Val)
		currentnode = currentnode.Next
	}

	s += strconv.Itoa(currentnode.Val)

	// convert binary string to base 10 representation
	decimalValue, _ := strconv.ParseInt(s, 2, 32)
	return int(decimalValue)
}

func printLinkedList(head *ListNode) {
	var currentNode *ListNode
	currentNode = head

	for currentNode.Next != nil {
		fmt.Println(currentNode.Val)
		currentNode = currentNode.Next
	}

	fmt.Println(currentNode.Val) // just so you don't miss the last element
}

func main() {
	// create 3 nodes
	head := ListNode{Val: 1}
	node2 := ListNode{Val: 0}
	node3 := ListNode{Val: 1}

	// add pointers to string them together into a linked list
	head.Next = &node2
	node2.Next = &node3

	fmt.Println(head)
	fmt.Println(node2)
	fmt.Println(node3)

	printLinkedList(&head)
	fmt.Println(getDecimalValue(&head))
}
