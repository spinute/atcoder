package main

import "fmt"

func max(ns []int) int {
	m := ns[0]
	for _, v := range ns {
		if v > m {
			m = v
		}
	}
	return m
}

func main() {
	var n int
	fmt.Scan(&n)
	ns := make([]int, n)
	for i := range ns {
		fmt.Scan(&ns[i])
	}

	sum := 0
	for _, v := range ns {
		sum += v
	}

	m := max(ns)

	if sum-m > m {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
