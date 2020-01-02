package main

import "fmt"

func is_square(n int) bool {
	for i := 0; i*i <= n; i++ {
		if i*i == n {
			return true
		}
	}
	return false
}

func concat(a, b int) int {
	ofs := 10
	for ofs <= b {
		ofs *= 10
	}
	return a*ofs + b
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	n := concat(a, b)

	if is_square(n) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
