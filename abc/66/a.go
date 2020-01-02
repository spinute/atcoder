package main

import (
	"fmt"
	"sort"
)

func main() {
	var ps [3]int
	for i := range ps {
		var p int
		fmt.Scan(&p)
		ps[i] = p
	}
	sort.Ints(ps[:])
	fmt.Println(ps[0] + ps[1])
}
