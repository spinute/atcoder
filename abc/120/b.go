package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)

	count := 0
	i := min(a, b)
	for {
		if a%i == 0 && b%i == 0 {
			count++
		}
		if count == k {
			fmt.Println(i)
			return
		}
		i--
	}
}
