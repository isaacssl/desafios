package main

import (
	"fmt"
)

func checkValidString(s string) bool {
	minOpen := 0
	maxOpen := 0

	for _, v := range s {
		switch v {
		case 40:
			minOpen++
			maxOpen++
		case 41:
			minOpen = bigger(minOpen-1, 0)
			maxOpen--
			if maxOpen < 0 {
				return false
			}
		default:
			minOpen = max(minOpen-1, 0)
			maxOpen++
		}
	}

	return minOpen == 0
}

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(checkValidString("(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"))
}
