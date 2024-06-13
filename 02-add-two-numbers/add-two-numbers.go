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
	if r+a+b >= 10 {
		return 1, r + a + b - 10
	}
	return 0, a + b + r
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head ListNode
	var v, r int
	stop := false
	r = 0

	curr := &head
	v1 := l1
	v2 := l2
	a1 := v1.Val
	a2 := v2.Val

	for !stop {
		r, v = sumOperation(r, a1, a2)
		curr.Val = v

		if v1.Next != nil || v2.Next != nil || r != 0 {
			var newNode ListNode
			curr.Next = &newNode
			curr = curr.Next

			if v1.Next != nil {
				v1 = v1.Next
				a1 = v1.Val
			} else {
				a1 = 0
			}
			if v2.Next != nil {
				v2 = v2.Next
				a2 = v2.Val
			} else {
				a2 = 0
			}
		} else {
			stop = true
		}

	}

	return &head
}

func main() {

	fmt.Println("start")

	arr := []int{2, 4, 3}
	arr2 := []int{5, 6, 4}

	head1 := createALL(arr)
	head2 := createALL(arr2)

	head := addTwoNumbers(head1, head2)

	printALL(head1, head2, head)

}
