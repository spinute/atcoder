package main

import "fmt"

func main() {
	var a, b, n int
	fmt.Scan(&a, &b, &n)
	for i := n; ; i++ {
		if i%a == 0 && i%b == 0 {
			fmt.Println(i)
			break
		}
	}
}
