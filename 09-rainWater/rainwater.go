package main

import "fmt"

func maxValue(vetor []int) int {
	var vMax int
	for _, v := range vetor {
		if v > vMax {
			vMax = v
		}
	}
	return vMax
}

func checkLine(xo int, xf int, ve []int) (int, int) {
	var curr, prev, count, temp int
	for i := xo; i < xf; i++ {
		curr = ve[i+1]
		prev = ve[1]
		if curr < prev {
			temp++
		}
		if curr > prev {
			count += temp
		}
	}

	return 0, 0
}

func qualquercoisa(vetor []int) int {

	var dx, dy int
	dx = len(vetor)
	dy = maxValue(vetor)

	var matriz [][]int

	if dx >= 1 && dy >= 1 {
		matriz[dx][dy] = 1

		fmt.Println(matriz)

	}

	return 0
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(height)
	i := qualquercoisa(height)
	fmt.Print(i)
}
