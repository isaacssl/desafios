package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, v := range nums {
		if j, ok := numMap[target-v]; ok && i != j {
			return []int{i, j}
		}
		numMap[v] = i
	}
	return nil // No solution found
}

func main() {
	Nums := []int{3, 50, 25, 12, 2, 4}
	Target := 5

	fmt.Println(twoSum(Nums, Target))
}
