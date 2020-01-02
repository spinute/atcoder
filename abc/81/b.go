package main

import "fmt"

func num_div(n int) int {
	if n%2 == 0 {
		return num_div(n/2) + 1
	} else {
		return 0
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}

	min := num_div(as[0])
	for _, a := range as {
		if nd := num_div(a); nd < min {
			min = nd
		}
	}
	fmt.Println(min)
}
