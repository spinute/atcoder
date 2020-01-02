package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)
	ts := make([]int, n)
	for i := range ts {
		fmt.Scan(&ts[i])
	}
	fmt.Scan(&m)
	ps := make([]int, m)
	xs := make([]int, m)
	for i := range ps {
		fmt.Scan(&ps[i])
		fmt.Scan(&xs[i])
	}

	for i, t := range ps {
		sum := 0
		for j := range ts {
			if j == t-1 {
				sum += xs[i]
			} else {
				sum += ts[j]
			}
		}
		fmt.Println(sum)
	}
}
