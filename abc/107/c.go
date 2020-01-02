package main

import "fmt"

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dist(r, l int) int {
	return min(diff(0, r)+diff(r, l), diff(0, l)+diff(l, r))
}

func main() {
	var N, K int
	fmt.Scan(&N, &K)
	xs := make([]int, N)
	for i := range xs {
		fmt.Scan(&xs[i])
	}

	min := dist(xs[0], xs[K-1])
	for i := 0; i+K-1 < len(xs); i++ {
		l := xs[i]
		r := xs[i+K-1]
		if v := dist(l, r); v < min {
			min = v
		}
	}
	fmt.Println(min)
}
