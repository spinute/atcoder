package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	var n, m, x int
	fmt.Scan(&n, &m, &x)
	as := make([]int, m)
	for i := range as {
		fmt.Scan(&as[i])
	}

	var cost_r, cost_l int
	for _, a := range as {
		if a < x {
			cost_r++
		} else {
			cost_l++
		}
	}

	fmt.Println(min(cost_l, cost_r))
}
