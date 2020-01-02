package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ls := make([]int, n)
	rs := make([]int, n)
	for i := range ls {
		fmt.Scan(&ls[i], &rs[i])
	}

	sum := 0
	for i := range ls {
		sum += rs[i] - ls[i] + 1
	}
	fmt.Println(sum)
}
