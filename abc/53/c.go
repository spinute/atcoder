package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	ans := x / 11 * 2
	if v := x % 11; v > 0 {
		ans++
		if v > 6 {
			ans++
		}
	}
	fmt.Println(ans)
}
