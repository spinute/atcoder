package main

import "fmt"

func min(vs []int) int {
	min_v := vs[0]
	for _, v := range vs {
		if v < min_v {
			min_v = v
		}
	}
	return min_v
}

func max(vs []int) int {
	max_v := vs[0]
	for _, v := range vs {
		if v > max_v {
			max_v = v
		}
	}
	return max_v
}

func main() {
	var n int
	fmt.Scan(&n)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}
	fmt.Println(max(as) - min(as))
}
