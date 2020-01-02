package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	var counter [10]int
	counter[a]++
	counter[b]++
	counter[c]++

	if counter[5] == 2 && counter[7] == 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
