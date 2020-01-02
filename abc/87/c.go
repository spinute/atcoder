package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var as [2][]int
	for i := range as {
		as[i] = make([]int, n)
		for j := range as[i] {
			fmt.Scan(&as[i][j])
		}
	}

	max := 0
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j <= i; j++ {
			sum += as[0][j]
		}
		for k := i; k < n; k++ {
			sum += as[1][k]
		}
		if sum > max {
			max = sum
		}
	}
	fmt.Println(max)
}
