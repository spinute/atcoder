package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	ans := ""
	for N != 0 {
		if N%2 == 0 {
			ans = "0" + ans
		} else {
			N--
			ans = "1" + ans
		}
		N /= -2
	}

	if ans == "" {
		ans = "0"
	}

	fmt.Println(ans)
}
