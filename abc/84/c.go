package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cs := make([]int, n-1)
	ss := make([]int, n-1)
	fs := make([]int, n-1)
	for i := range cs {
		fmt.Scan(&cs[i], &ss[i], &fs[i])
	}

	for i := range cs {
		t := 0
		for j := i; j < n-1; j++ {
			if ss[j] > t {
				t = ss[j]
			} else {
				t = ((t-1)/fs[j] + 1) * fs[j]
			}
			t += cs[j]
		}
		fmt.Println(t)
	}
	fmt.Println(0)
}
