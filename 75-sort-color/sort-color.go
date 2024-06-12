package main

import "fmt"

func sortColors(nums []int) {
	nomorechange := false
	for !nomorechange {
		nomorechange = true
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nomorechange = false
				curr := nums[i]
				nums[i] = nums[i+1]
				nums[i+1] = curr
			}
		}
	}
}

func main() {
	var nums []int
	fmt.Println("start")

	nums = []int{2, 0, 2, 1, 1, 0}

	sortColors(nums)

	fmt.Println(nums)

}
