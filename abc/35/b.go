package main

import "fmt"

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	var s string
	var t int
	fmt.Scan(&s, &t)

	counter := make(map[rune]int)
	for _, r := range s {
		counter[r]++
	}

	if t == 1 {
		fmt.Println(diff(counter['R'], counter['L']) +
			diff(counter['U'], counter['D']) + counter['?'])
	} else {
		v := diff(counter['R'], counter['L']) + diff(counter['U'], counter['D'])
		u := counter['?']
		if v > u {
			fmt.Println(v - u)
		} else {
			fmt.Println(diff(v, u) % 2)
		}
	}
}
