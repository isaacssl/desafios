package main

func shiftLine(tickets *[]int, target *[]int) {
	tmptikets := *tickets
	head := tmptikets[0]
	headtrg := target[0]
	tmptikets = append(tmptikets, head)
	tmptikets = tmptikets[1:]
	tmptarget := append(*target, headtrg)
	tickets = &tmptikets
	target = &tmptarget[1:]
}

func outline(tickets *[]int, target *[]int) {
	tickets = &tickets[1:]
	target = &target[1:]
}

func timeRequiredToBuy(tickets []int, k int) int {

	var c, time int
	lenLine := len(tickets)

	target := make([]int, len(tickets))
	target[k] = 1

	for {
		tickets[0]--
		if tickets[0] != 0 {
			shiftLine(&tickets, &target)
		}
		if tickets[0] == 0 {
			if target[0] == 0 {
				return time
			}
			outline(&tickets, &target)
		}
		time++
		c++
		if c == lenLine {
			c = 0
			lenLine = len(tickets)
		}

	}

}

func main() {

	ticket := []int{2, 3, 2}
	k := 2

}
