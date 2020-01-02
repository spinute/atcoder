package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	counter := make(map[rune]int)
	for _, r := range s {
		counter[r]++
	}
	fmt.Println(counter['A'], counter['B'], counter['C'],
		counter['D'], counter['E'], counter['F'])
}
