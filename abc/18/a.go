package main

import "fmt"

func point(a, b int) int {
	if a > b {
		return 0
	}
	return 1
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(point(a, b) + point(a, c) + 1)
	fmt.Println(point(b, a) + point(b, c) + 1)
	fmt.Println(point(c, b) + point(c, a) + 1)
}
