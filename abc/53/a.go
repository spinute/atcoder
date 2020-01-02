package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a < 1200 {
		fmt.Println("ABC")
	} else {
		fmt.Println("ARC")
	}
}
