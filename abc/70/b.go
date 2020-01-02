package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if v := min(b, d) - max(a, c); v > 0 {
		fmt.Println(v)
	} else {
		fmt.Println(0)
	}
}
