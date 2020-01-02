package main

import "fmt"

func f(a int) int {
	ans := 0
	switch a % 4 {
	case 0:
		ans = a
	case 1:
		ans = 1
	case 2:
		ans = a/4*4 + 3
	case 3:
		ans = 0
	}
	return ans
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a == 0 {
		fmt.Println(f(b))
	} else {
		fmt.Println(f(b) ^ f(a-1))
	}
}
