package main

import "fmt"

func cube(a int) int {
	return a * a * a
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	sum := cube(n / k)
	if k%2 == 0 {
		sum += cube(n/(k/2) - n/k)
	}
	fmt.Println(sum)
}
