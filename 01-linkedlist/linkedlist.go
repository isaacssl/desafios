package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type Linkedlist struct {
	Head *Node
	Tail *Node
	Len  int
}

// Insert a new node in Linkedlist Tail
func (l *Linkedlist) Insert(val int) {

	curr := &Node{Val: val}

	if l.Head != nil {
		l.Tail.Next = curr
	} else {
		l.Head = curr
	}

	l.Tail = curr
	l.Len++
}

// Find a Node by value - Only first Node
func (l *Linkedlist) Find(val int) (int, **Node) {
	var curr *Node
	var pointer **Node

	if l.Len == 0 {
		return -1, nil
	}

	curr = l.Head
	pointer = &l.Head

	for i := range l.Len {
		if curr.Val == val {
			return i, pointer
		}
		pointer = &curr.Next
		curr = curr.Next
	}
	return -1, nil
}

// Find all Nodes by value
func (l *Linkedlist) FindAll(val int) ([]int, []**Node) {

	var ix []int
	var pointers []**Node
	curr := l.Head
	po := &l.Head

	for i := range l.Len {
		if curr.Val == val {
			pointers = append(pointers, po)
			ix = append(ix, i)
		}
		po = &curr.Next
		curr = curr.Next
	}
	return ix, pointers
}

// Delete the first Node with value "val" or retunr a erro
func (l *Linkedlist) Del(val int) error {

	if l.Len == 0 {
		return fmt.Errorf("This LinkedList is empty")
	}

	if l.Head.Val == val {
		l.Head = l.Head.Next
		l.Len--
		return nil
	}

	var prev, curr *Node = l.Head, l.Head.Next

	for i := range l.Len - 1 {
		if curr.Val == val {
			prev.Next = curr.Next
			l.Len--
			break
		}
		prev = curr
		curr = curr.Next
		i++
	}
	return nil
}

// Delete all Nodes with value "val" or retunr a erro
func (l *Linkedlist) DellAll(val int) error {

	if l.Len == 0 {
		return fmt.Errorf("This LinkedList is empty")
	}

	curr := l.Head

	for curr.Val == val {
		l.Head = curr.Next
		curr = curr.Next
		l.Len--
	}

	if l.Head == nil {
		l.Tail = nil
		return nil
	}

	prev := l.Head
	curr = l.Head.Next
	intera := 1

	for curr != nil {
		for curr.Val == val {
			prev.Next = curr.Next
			curr = curr.Next
			l.Len--
		}
		prev = prev.Next
		curr = curr.Next
		intera++
	}

	return nil
}

// Create a array with the linkedlist values
func (l *Linkedlist) ToArray() []int {

	var vetor []int
	var curr *Node

	if l.Len == 0 {
		return nil
	}

	curr = l.Head

	for i := range l.Len {
		vetor = append(vetor, curr.Val)
		i++
		curr = curr.Next
	}

	return vetor
}

func main() {

	var LL Linkedlist
	vetor := []int{3, 3, 3, 5, 3, 8, 9, 8, 3, 6, 3, 3, 3, 3}

	fmt.Println(LL)

	for _, val := range vetor {
		LL.Insert(val)
	}

	fmt.Println(LL.ToArray())

	fmt.Printf("O ponteiro de LL é: %v\n", &LL.Head)
	w, five := LL.Find(3)
	fmt.Printf("O indice é %d e o ponteiro é %v\n", w, five)
	r, tree := LL.FindAll(3)
	fmt.Printf("O indice é %d e o ponteiro é %v\n", r, tree)

	LL.DellAll(3)
	fmt.Println(LL.ToArray())

}
