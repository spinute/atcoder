package main

import "fmt"

const c = 1000000007

func fact(n int64) int64 {
	if n == 1 {
		return 1
	}
	return n * fact(n-1) % c
}

func solve(n, m int64) int64 {
	switch n {
	case m:
		return fact(n) * fact(m) * 2 % c
	case m + 1, m - 1:
		return fact(n) * fact(m) % c
	}
	return 0
}

func main() {
	var n, m int64
	fmt.Scan(&n, &m)
	fmt.Println(solve(n, m))
}
