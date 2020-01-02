package main

import (
	"fmt"
	"sort"
)

func median(vs []int) int {
	sort.Ints(vs)
	return vs[len(vs)/2]
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func score(vs []int, b int) int {
	sum := 0
	for i, v := range vs {
		sum += abs(v - (b + i + 1))
	}
	return sum
}

func main() {
	var N int
	fmt.Scan(&N)
	as := make([]int, N)
	for i := range as {
		fmt.Scan(&as[i])
	}

	bs := make([]int, N)
	for i := range bs {
		bs[i] = as[i] - i - 1
	}
	b := median(bs)

	fmt.Println(score(as, b))
}
