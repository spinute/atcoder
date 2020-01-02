package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)

	if a == "a" {
		fmt.Println(-1)
	} else {
		fmt.Println("a")
	}
}
