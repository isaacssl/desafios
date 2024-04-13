package main

import (
	"fmt"
)

func goback(vector *[]int) {
	v := *vector
	v = append(v, v[0])
	*vector = v[1:]
}

func coupeLaTete(vector *[]int) {
	v := *vector
	*vector = v[1:]
}

func countStudents(students []int, sandwiches []int) int {
	rest := len(sandwiches)
	stdtCount := len(students)

	for len(sandwiches) > 0 {
		switch {
		case sandwiches[0] == students[0]:
			coupeLaTete(&students)
			coupeLaTete(&sandwiches)
			stdtCount = len(students)
			rest--
		case stdtCount == 0:
			return rest
		default:
			goback(&students)
			stdtCount--
		}
	}
	return 0
}

func main() {
	students := []int{1, 1, 1, 0, 0, 1}
	sandwiches := []int{1, 0, 0, 0, 1, 1}

	fmt.Println(countStudents(students, sandwiches))
}
