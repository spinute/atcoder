package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	i := 1
	for {
		if i*i*i*i == x {
			fmt.Println(i)
			break
		}
		i++
	}
}
