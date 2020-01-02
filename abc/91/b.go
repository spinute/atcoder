package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)
	ss := make([]string, n)
	for i := range ss {
		fmt.Scan(&ss[i])
	}
	fmt.Scan(&m)
	ts := make([]string, m)
	for i := range ts {
		fmt.Scan(&ts[i])
	}

	earns := make(map[string]int)
	for _, s := range ss {
		earns[s]++
	}
	for _, t := range ts {
		earns[t]--
	}

	max_earn := 0
	for _, v := range earns {
		if v > max_earn {
			max_earn = v
		}
	}
	fmt.Println(max_earn)
}
