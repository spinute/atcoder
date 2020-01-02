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
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(c / min(a, b))
}
