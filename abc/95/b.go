package main

import "fmt"

func min(ms []int) int {
	min_value := ms[0]
	for _, v := range ms {
		if v < min_value {
			min_value = v
		}
	}
	return min_value
}

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	ms := make([]int, n)
	for i := range ms {
		fmt.Scan(&ms[i])
		x -= ms[i]
	}
	fmt.Println(n + x/min(ms))
}
