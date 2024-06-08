package main

import "fmt"

type treeNode struct{
	Val int
	Root *treeNode
	leftNode *treeNode
	rightNode *treeNode
}

type Tree struct {
	root *treeNode
	deep int
	last *treeNode
	next byte
}

func insertNode(v int, T *Tree) {
	newnode := treeNode{Val: v}

	var

	
}

func creatTree(v []int) *[]int {

}

func main() {
	root := []int{1, 2, 3}
	fmt.Println(root)
}
