package main

import "fmt"

func main() {
	var n, a int
	fmt.Scan(&n, &a)
	xs := make([]int, n)
	for i := range xs {
		fmt.Scan(&xs[i])
	}
}
