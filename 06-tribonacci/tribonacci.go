package main

import "fmt"

func recursive(i *int, stop *int, v *int) int {
	//ti := *i + 1
	*i += 1
	*v += 1

	fmt.Println(*v, v)

	if *i != *stop {
		recursive(i, stop, v)
	}
	return 2
}

func triIntegrator(sum *int, i *int, n *int, v *[]int) {
	vtemp := *v
	T := vtemp[0] + vtemp[1] + vtemp[2]
	*sum = T
	if *i < *n {
		vtemp = append(vtemp, T)
		vtemp = vtemp[1:]
		v = &vtemp
		*i += 1
		triIntegrator(sum, i, n, v)
	}
}

func tribonacci(n int) int {
	var sum int
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	default:
		v := []int{0, 1, 1}
		i := 3
		triIntegrator(&sum, &i, &n, &v)
	}
	return sum
}

func main() {
	//i := 0
	//v := 0
	//stop := 10
	//recursive(&i, &stop, &v)

	n := 25
	Hu := tribonacci(n)
	fmt.Println(Hu)

}
