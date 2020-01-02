package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	power := int64(1)
	for i := 0; i < n; i++ {
		power = power * int64(i+1) % 1000000007
	}
	fmt.Println(power)
}
