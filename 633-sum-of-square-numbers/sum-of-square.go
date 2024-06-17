package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	for a := 0; a*a <= c; a++ {
		bSquared := c - a*a
		b := int(math.Sqrt(float64(bSquared)))
		if b*b == bSquared {
			return true
		}
	}
	return false
}

func main() {
	n := 0

	etAlors := judgeSquareSum(n)

	fmt.Println(etAlors)
}
