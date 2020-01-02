package main

import "fmt"

func pow100(d int) int {
	if d == 0 {
		return 1
	} else {
		return 100 * pow100(d-1)
	}
}

func main() {
	var d, n int
	fmt.Scan(&d, &n)
	if n != 100 {
		fmt.Println(pow100(d) * n)
	} else {
		fmt.Println(pow100(d) * 101)
	}
}
