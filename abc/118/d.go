package main

import (
	"fmt"
	"sort"
)

var cs = []int{9, 2, 5, 5, 4, 5, 6, 3, 7, 6}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	as := make([]int, m)
	for i := range as {
		fmt.Scan(&as[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(as)))

	reachable := make([]bool, n+1)
	digits := make([]int, n+1)
	reachable[0] = true

	for i := range reachable {
		for _, a := range as {
			j := i - cs[a]
			if j >= 0 && reachable[j] {
				if v := digits[j] + 1; v > digits[i] {
					digits[i] = v
					reachable[i] = true
				}
			}
		}
	}

	for i := n; i != 0; {
		for _, a := range as {
			j := i - cs[a]
			if j >= 0 && reachable[j] && digits[j]+1 == digits[i] {
				fmt.Print(a)
				i = j
				break
			}
		}
	}
	fmt.Println("")
}
