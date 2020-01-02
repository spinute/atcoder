package main

import "fmt"

func main() {
	var x, y [4]int
	fmt.Scan(&x[0], &y[0], &x[1], &y[1])
	x[2] = y[0] - y[1] + x[1]
	y[2] = -x[0] + x[1] + y[1]
	x[3] = y[1] - y[2] + x[2]
	y[3] = -x[1] + x[2] + y[2]
	fmt.Println(x[2], y[2], x[3], y[3])
}
