package main

import "fmt"

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func solve(as []int, next int) int {
	sum := 0
	count := 0
	for _, a := range as {
		sum += a
		if next*sum <= 0 {
			count += diff(sum, next)
			sum = next
		}
		next *= -1
	}
	return count
}

func main() {
	var n int
	fmt.Scan(&n)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}

	fmt.Println(min(solve(as, 1), solve(as, -1)))
}
