package main

import (
	"fmt"
	"sort"
)

func sum(vs []int) int {
	s := 0
	for _, v := range vs {
		s += v
	}
	return s
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	xs := make([]int, m)
	for i := range xs {
		fmt.Scan(&xs[i])
	}
	sort.Ints(xs)

	ds := make([]int, m-1)
	for i := range ds {
		ds[i] = xs[i+1] - xs[i]
	}

	sort.Sort(sort.Reverse(sort.IntSlice(ds)))

	if n > m {
		fmt.Println(0)
	} else {
		fmt.Println(sum(ds[n-1:]))
	}
}
