package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scan(&n)
	as := make([]int, max(n, 3))
	as[0] = 0
	as[1] = 0
	as[2] = 1
	for i := 3; i < n; i++ {
		as[i] = (as[i-1] + as[i-2] + as[i-3]) % 10007
	}
	fmt.Println(as[n-1])
}
