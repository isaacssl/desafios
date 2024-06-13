package main

import "fmt"

func adsSubtract(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return b - a
}

func minMovesToSeat(seats []int, students []int) int {

	hasChange := true
	var distance int

	for hasChange {
		distance = 0
		hasChange = false
		for i := 0; i < len(seats)-1; i++ {
			if seats[i] > seats[i+1] {
				hasChange = true
				temp := seats[i]
				seats[i] = seats[i+1]
				seats[i+1] = temp
			}
			if students[i] > students[i+1] {
				hasChange = true
				temp := students[i]
				students[i] = students[i+1]
				students[i+1] = temp
			}
			distance += adsSubtract(seats[i], students[i])
		}
		distance += adsSubtract(seats[len(seats)-1], students[len(students)-1])
	}
	return distance
}

func main() {
	seats := []int{2, 2, 6, 6}
	students := []int{1, 3, 2, 6}

	distance := minMovesToSeat(seats, students)

	fmt.Println(distance)

}
