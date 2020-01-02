package main

import "fmt"

func main() {
	var x, a, b int
	fmt.Scan(&x, &a, &b)
	if b <= a {
		fmt.Println("delicious")
	} else if b <= a+x {
		fmt.Println("safe")
	} else {
		fmt.Println("dangerous")
	}
}
