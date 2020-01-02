package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(max((a+1)*b, a*(b+1)))
}
