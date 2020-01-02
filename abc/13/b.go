package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	d := diff(a, b)
	fmt.Println(min(d, 10-d))
}
