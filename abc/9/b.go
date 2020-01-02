package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}

	sort.Ints(as)
	max := as[len(as)-1]
	for i := len(as) - 1; i >= 0; i-- {
		if as[i] == max {
			continue
		}
		fmt.Println(as[i])
		break
	}
}
