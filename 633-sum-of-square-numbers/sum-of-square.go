package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	cfloat := float64(c)
	var a float64
	max := cfloat * math.Sqrt(2)

	for a = 0; a <= max; a++ {
		b := math.Sqrt(cfloat - (a * a))
		if math.Mod(b, 1) == 0 {
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
