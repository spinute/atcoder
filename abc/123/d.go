package main

import "fmt"

func main() {
	var x, y, z, k int
	fmt.Scan(&x, &y, &z, &k)
	as := make([]int, x)
	bs := make([]int, y)
	cs := make([]int, z)
	for i := range as {
		fmt.Scan(&as[i])
	}
	for i := range bs {
		fmt.Scan(&bs[i])
	}
	for i := range cs {
		fmt.Scan(&cs[i])
	}
}
