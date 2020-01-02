package main

import "fmt"

func abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if abs(a-c) <= d {
		fmt.Println("Yes")
	} else if abs(a-b) <= d && abs(c-b) <= d {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
