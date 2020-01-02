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
	score := 0
	for i := 0; i < len(as); i++ {
		if i%2 == 0 {
			score += as[len(as)-1-i]
		} else {
			score -= as[len(as)-1-i]
		}
	}
	fmt.Println(score)
}
