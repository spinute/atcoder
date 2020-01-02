package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	cost := 0
	for i := 1; i < len(s); i++ {
		if s[i] == 'E' {
			cost++
		}
	}
	min := cost
	for i := 1; i < len(s); i++ {
		if s[i-1] == 'W' {
			cost++
		}
		if s[i] == 'E' {
			cost--
		}
		if cost < min {
			min = cost
		}
	}

	fmt.Println(min)
}
