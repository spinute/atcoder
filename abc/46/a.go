package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a == b && b == c {
		fmt.Println(1)
	} else if a == b || b == c || c == a {
		fmt.Println(2)
	} else {
		fmt.Println(3)
	}
}
