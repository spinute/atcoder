package main

import (
	"fmt"
	"sort"
)

func pow2(n int) int {
	if n == 0 {
		return 1
	} else {
		return pow2(n-1) * 2
	}
}

func main() {
	var k int
	vs := make([]int, 3)
	fmt.Scan(&vs[0], &vs[1], &vs[2], &k)
	sort.Ints(vs)
	fmt.Println(vs[0] + vs[1] + vs[2]*pow2(k))
}
