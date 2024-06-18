package main

import (
	"fmt"
	"sort"
)

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	relation := make(map[int]int)

	for i, v := range difficulty {
		if relation[v] < profit[i] {
			relation[v] = profit[i]
		}
	}
	sort.Ints(difficulty)
	total := 0

	for _, w := range worker {
		bestBudge := 0
		for _, d := range difficulty {
			if w < d {
				break
			}
			if relation[d] > bestBudge {
				bestBudge = relation[d]
			}

		}
		total += bestBudge

	}
	return total
}

func main() {

	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	difficulty := []int{68, 35, 52, 47, 86}
	profit := []int{67, 17, 1, 81, 3}
	worker := []int{92, 10, 85, 84, 82}

	total := maxProfitAssignment(difficulty, profit, worker)

	fmt.Println(total)

	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	difficulty = []int{66, 1, 28, 73, 53, 35, 45, 60, 100, 44, 59, 94, 27, 88, 7, 18, 83, 18, 72, 63}
	profit = []int{66, 20, 84, 81, 56, 40, 37, 82, 53, 45, 43, 96, 67, 27, 12, 54, 98, 19, 47, 77}
	worker = []int{61, 33, 68, 38, 63, 45, 1, 10, 53, 23, 66, 70, 14, 51, 94, 18, 28, 78, 100, 16}

	total = maxProfitAssignment(difficulty, profit, worker)

	fmt.Println(total)

}
