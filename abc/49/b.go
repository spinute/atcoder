package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	cs := make([]string, h)
	for i := range cs {
		fmt.Scan(&cs[i])
	}

	for _, s := range cs {
		fmt.Println(s)
		fmt.Println(s)
	}
}
