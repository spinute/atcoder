package main

import "fmt"

func group(a int) int {
	switch a {
	case 1, 3, 5, 7, 8, 10, 12:
		return 1
	case 4, 6, 9, 11:
		return 2
	case 2:
		return 3
	}
	return -1
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if group(a) == group(b) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
