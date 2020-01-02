package main

import "fmt"

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	var n int
	fmt.Scan(&n)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}

	sum := 0
	for _, a := range as {
		sum += a
	}

	min := diff(as[0], sum-as[0])
	su := 0
	for i := 0; i < len(as)-1; i++ {
		su += as[i]
		if v := diff(su, sum-su); v < min {
			min = v
		}
	}

	fmt.Println(min)
}
