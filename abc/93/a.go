package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	counter := make(map[rune]int)

	for _, c := range s {
		counter[c] += 1
	}

	if counter['a'] == 1 && counter['b'] == 1 && counter['c'] == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
