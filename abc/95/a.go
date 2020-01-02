package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	price := 700
	for _, c := range s {
		if c == 'o' {
			price += 100
		}
	}
	fmt.Println(price)
}
