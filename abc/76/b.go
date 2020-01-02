package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	num := 1
	for i := 0; i < n; i++ {
		if num*2 < num+k {
			num *= 2
		} else {
			num += k
		}
	}
	fmt.Println(num)
}
