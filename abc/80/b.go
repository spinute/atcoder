package main

import "fmt"

func sum_digits(n int) int {
	if n < 10 {
		return n
	} else {
		return n%10 + sum_digits(n/10)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	if n%sum_digits(n) == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
