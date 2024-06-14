package main

import (
	"fmt"
)

func minIncrementForUnique(nums []int) int {
	numsMap := make(map[int]bool)
	move := 0

	for _, v := range nums {
		for numsMap[v] {
			v++
			move++
		}
		numsMap[v] = true
	}
	return move
}

func main() {
	num0 := []int{1, 2, 2}
	num1 := []int{3, 2, 1, 2, 1, 7}

	ans0 := minIncrementForUnique(num0)
	ans1 := minIncrementForUnique(num1)

	fmt.Println(ans0, ans1)
}
