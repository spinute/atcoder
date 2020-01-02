package main

import (
	"fmt"
	"sort"
)

func main() {
	vs := make([]int, 3)
	fmt.Scan(&vs[0], &vs[1], &vs[2])
	sort.Ints(vs)

	res := vs[2] - vs[1] + vs[2] - vs[0]
	if res%2 == 0 {
		fmt.Println(res / 2)
	} else {
		fmt.Println((res + 3) / 2)
	}
}
