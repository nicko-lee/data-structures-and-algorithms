package main

// //   Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// type LinkedList struct {
// 	Head *ListNode
// }

// func (ll LinkedList) storeData(dataCollection []int) LinkedList {
// 	// create the head of the linked list
// 	var headNode LinkedList

// 	// grab each item in data collection and put it in a ListNode
// 	for i, v := range dataCollection {
// 		// create two new listNode structs and keep track of their pointers
// 		pointerToCurrentNode := &ListNode{}
// 		pointerToNextNode := &ListNode{}

// 		// first check for head element - all we need to keep track of is the head
// 		if i == 0 {
// 			headNode.Head = pointerToCurrentNode
// 		}

// 		// populate value of new listNode including a pointer to the next node else you break the chain
// 		*pointerToCurrentNode.Val = v
// 		*pointerToCurrentNode.Next = pointerToNextNode

// 		// check for last element of linked list - and remove the pointer to signify list ends here when there is a null pointer
// 		if i == len(dataCollection-1) {
// 			*pointerToCurrentNode.Next = nil
// 		}
// 	}

// 	return headNode
// }

// func (ll LinkedList) getData(head LinkedList) []int {
// 	// create a variable to store results of list traversal
// 	var storeOutput []int
// 	var nextNode *ListNode

// 	// grab first element's value and append to storeOutput
// 	storeOutput = append(storeOutput, *head.Head.Val)
// 	nextNode = *head.Head.Next

// 	for nextNode != nil {
// 		// grab next element
// 	}

// 	for i, v := range dataCollection {

// 	}
// }

// func getDecimalValue(head *ListNode) int {
// 	return 1
// }

// // func main() {
// // 	var newLinkedList LinkedList
// // 	input := newLinkedList

// // 	fmt.Println(input)
// // }

// func getDecimalValue(head *ListNode) int {
// 	sum := 0
// 	for head != nil {
// 		sum = sum*2 + head.Val
// 		head = head.Next
// 	}
// 	return sum
// }
