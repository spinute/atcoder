package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	max := 0
	cnt := 0
	for _, r := range s {
		if r == 'I' {
			cnt++
		} else {
			cnt--
		}

		if cnt > max {
			max = cnt
		}
	}

	fmt.Println(max)
}
