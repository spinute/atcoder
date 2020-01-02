package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ss := make([]string, n)
	for i := range ss {
		fmt.Scan(&ss[i])
	}

	m := make(map[string]int)
	for _, s := range ss {
		m[s]++
	}

	max := 0
	name := ""
	for k, v := range m {
		if v > max {
			name = k
			max = v
		}
	}
	fmt.Println(name)
}
