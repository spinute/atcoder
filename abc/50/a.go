package main

import "fmt"

func main() {
	var a, b int
	var o string
	fmt.Scan(&a, &o, &b)
	if o == "+" {
		fmt.Println(a + b)
	} else {
		fmt.Println(a - b)
	}
}
