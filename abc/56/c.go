package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	i := 1
	sum := 0
	for {
		sum += i
		if sum >= x {
			break
		}
		i++
	}
	fmt.Println(i)
}
