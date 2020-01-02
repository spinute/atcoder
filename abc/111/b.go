package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	for x := 111; x < 1000; x += 111 {
		if x >= a {
			fmt.Println(x)
			break
		}
	}
}
