```go
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// func to create a LinkedList
func createALL(arr []int) *ListNode {
	var head ListNode
	curr := &head
	curr.Val = arr[0]

	for i := 1; i < len(arr); i++ {
		newNode := ListNode{Val: arr[i]}
		curr.Next = &newNode
		curr = curr.Next
	}
	return &head
}

// func to print the ListNode values
func printLL(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("Value:%d <> Next:%v\n", curr.Val, &curr.Next)
		curr = curr.Next
	}
}

func printALL(h1, h2, h *ListNode) {
	curr1 := h1
	curr2 := h2
	curr := h
	var a, b int

	for curr != nil {
		if curr1 != nil {
			a = curr1.Val
			curr1 = curr1.Next
		} else {
			a = 0
		}
		if curr2 != nil {
			b = curr2.Val
			curr2 = curr2.Next
		} else {
			b = 0
		}
		fmt.Printf("%v + %v = %v\n", a, b, curr.Val)
		curr = curr.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func sumOperation(r, a, b int) (c, d int) {
	if r+a+b > 10 {
		return 1, r + a + b - 10
	}
	return 0, a + b + r
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head ListNode
	prev := &head

	// First Node
	curr1 := l1
	curr2 := l2

	r, v := sumOperation(0, curr1.Val, curr2.Val)

	prev.Val = v

	// Next Nodes
	for curr1.Next != nil && curr2.Next != nil {

		var newNode ListNode

		prev.Next = &newNode
		curr1 = curr1.Next
		curr2 = curr2.Next
		prev = prev.Next

		r, v = sumOperation(r, curr1.Val, curr2.Val)
		prev.Val = v
	}

	for curr1.Next != nil {
		if r != 0 {
			var newNode ListNode

			prev.Next = &newNode
			curr1 = curr1.Next
			r, v = sumOperation(r, curr1.Val, 0)
			prev.Val = v
		} else {
			prev.Next = curr1.Next
			break
		}
	}

	for curr2.Next != nil {
		if r != 0 {
			var newNode ListNode

			prev.Next = &newNode
			curr2 = curr1.Next
			r, v = sumOperation(r, curr2.Val, 0)
			prev.Val = v
		} else {
			prev.Next = curr2.Next
			break
		}
	}
	return &head
}

func main() {

	fmt.Println("start")

	arr := []int{0, 9, 8, 9, 0, 9, 8}
	arr2 := []int{0, 9, 8, 9}

	head1 := createALL(arr)
	head2 := createALL(arr2)

	head := addTwoNumbers(head1, head2)

	printALL(head1, head2, head)

}
```
