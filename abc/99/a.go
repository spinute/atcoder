package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a < 1000 {
		fmt.Println("ABC")
	} else {
		fmt.Println("ABD")
	}
}
