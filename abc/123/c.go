package main

import "fmt"

func roundup(a, b int) int {
	if a%b == 0 {
		return a / b
	}
	return a/b + 1
}

func main() {
	var n int
	fmt.Scan(&n)
	xs := make([]int, 5)
	for i := range xs {
		fmt.Scan(&xs[i])
	}

	min := xs[0]
	for _, x := range xs {
		if x < min {
			min = x
		}
	}

	fmt.Println(4 + roundup(n, min))
}
