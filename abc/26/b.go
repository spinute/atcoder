package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	rs := make([]int, n)
	for i := range rs {
		fmt.Scan(&rs[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(rs)))
	sum := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			sum += rs[i] * rs[i]
		} else {
			sum -= rs[i] * rs[i]
		}
	}
	fmt.Println(float64(sum) * math.Pi)
}
