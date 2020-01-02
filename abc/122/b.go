package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	ans := 0
	cnt := 0
	for _, c := range s {
		if c == 'A' || c == 'C' || c == 'G' || c == 'T' {
			cnt++
		} else {
			cnt = 0
		}
		if cnt > ans {
			ans = cnt
		}
	}
	fmt.Println(ans)
}
