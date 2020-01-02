package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if u, v := a*d, b*c; u > v {
		fmt.Println("AOKI")
	} else if u == v {
		fmt.Println("DRAW")
	} else {
		fmt.Println("TAKAHASHI")
	}
}
