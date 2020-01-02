package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a > b {
		fmt.Println("Worse")
	} else {
		fmt.Println("Better")
	}
}
