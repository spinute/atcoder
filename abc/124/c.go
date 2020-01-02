package main

import (
	"fmt"
	"strconv"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var s string
	fmt.Scan(&s)

	ans1 := 0
	for i, a := range s {
		j, _ := strconv.Atoi(string(a))
		if i%2 != j {
			ans1++
		}
	}
	ans2 := 0
	for i, a := range s {
		j, _ := strconv.Atoi(string(a))
		if (i+1)%2 != j {
			ans2++
		}
	}

	fmt.Println(min(ans1, ans2))
}
