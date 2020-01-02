package main

import "fmt"

func check(s string) bool {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	for _, cnt := range m {
		if cnt%2 != 0 {
			return false
		}
	}
	return true
}

func main() {
	var s string
	fmt.Scan(&s)

	if check(s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
