package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve(n, d int) int {
	if n < 6 {
		return n + d
	}
	v1 := 1
	for v1*6 <= n {
		v1 *= 6
	}
	v2 := 1
	for v2*9 <= n {
		v2 *= 9
	}
	return min(solve(n-v1, d+1), solve(n-v2, d+1))
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solve(n, 0))
}
