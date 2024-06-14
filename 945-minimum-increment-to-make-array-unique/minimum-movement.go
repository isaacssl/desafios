package main

import "fmt"

func minIncrementForUnique(nums []int) int {
	hasChange := true
	move := 0

	for hasChange {
		hasChange = false
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				hasChange = true
				tmp := nums[i]
				nums[i] = nums[i+1]
				nums[i+1] = tmp
			}
			if nums[i] == nums[i+1] {
				hasChange = true
				nums[i+1]++
				move++
				break
			}
		}
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
