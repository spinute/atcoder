package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a < 1200 {
		fmt.Println("ABC")
	} else if a < 2800 {
		fmt.Println("ARC")
	} else {
		fmt.Println("AGC")
	}
}
