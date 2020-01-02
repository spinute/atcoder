package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	ln := make([]int, n)
	for i := range ln {
		fmt.Scan(&ln[i])
	}
	sort.Ints(ln)
	sum := 0
	for i := 0; i < k; i++ {
		sum += ln[n-1-i]
	}
	fmt.Println(sum)
}
