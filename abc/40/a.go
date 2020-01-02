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
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(min(b-1, a-b))
}
