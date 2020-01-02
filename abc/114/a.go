package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 7 || n == 5 || n == 3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
