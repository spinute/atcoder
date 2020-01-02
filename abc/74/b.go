package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	xs := make([]int, n)
	for i := range xs {
		fmt.Scan(&xs[i])
	}

	sum := 0
	for _, x := range xs {
		sum += 2 * min(x, k-x)
	}
	fmt.Println(sum)
}
