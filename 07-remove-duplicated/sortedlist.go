package main

import "fmt"

func removeDuplicates(nums *[]int) int {
	k := 1
	newNums := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			newNums = append(newNums, nums[i])
			k++
		}
	}
	nums = newNums
	return k
}

func main() {
	nums := []int{1, 1, 2}

	output := removeDuplicates(nums)

	fmt.Println(output)
}
