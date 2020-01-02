package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	prev := 'x'
	cnt := 0
	for _, c := range s {
		if prev != c {
			cnt++
		}
		prev = c
	}
	fmt.Println(cnt - 1)
}
