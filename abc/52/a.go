package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if v1, v2 := a*b, c*d; v1 > v2 {
		fmt.Println(v1)
	} else {
		fmt.Println(v2)
	}
}
