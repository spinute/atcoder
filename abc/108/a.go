package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println((a + 1) / 2 * (a / 2))
}