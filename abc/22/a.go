package main

import "fmt"

func main() {
	var n, s, t int
	fmt.Scan(&n, &s, &t)
	w := 0
	cnt := 0
	for i := 0; i < n; i++ {
		var df int
		fmt.Scan(&df)
		w += df
		if s <= w && w <= t {
			cnt++
		}
	}
	fmt.Println(cnt)
}
