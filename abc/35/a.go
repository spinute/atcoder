package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a*3 == b*4 {
		fmt.Println("4:3")
	} else {
		fmt.Println("16:9")
	}
}
