package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(min(a, b) + min(c, d))
}
