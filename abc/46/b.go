package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	ans := k
	for i := 1; i < n; i++ {
		ans *= k - 1
	}
	fmt.Println(ans)
}
