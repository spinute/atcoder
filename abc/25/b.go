package main

import "fmt"

func dist(a, b, d int) int {
	if d < a {
		return a
	} else if d > b {
		return b
	}
	return d
}

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	ss := make([]string, n)
	ds := make([]int, n)
	for i := range ss {
		fmt.Scan(&ss[i], &ds[i])
	}

	sum := 0
	for i, s := range ss {
		if s == "East" {
			sum += dist(a, b, ds[i])
		} else {
			sum -= dist(a, b, ds[i])
		}
	}

	if sum > 0 {
		fmt.Println("East", sum)
	} else if sum < 0 {
		fmt.Println("West", -sum)
	} else {
		fmt.Println(0)
	}
}
