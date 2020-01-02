package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 3; i++ {
		var s, e int
		fmt.Scan(&s, &e)
		sum += s / 10 * e
	}
	fmt.Println(sum)
}
