package main

import (
	"fmt"
	"strings"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	ss := make([]string, h)
	for i := range ss {
		fmt.Scan(&ss[i])
	}
	fmt.Println(strings.Repeat("#", w+2))
	for _, s := range ss {
		fmt.Printf("#%s#\n", s)
	}
	fmt.Println(strings.Repeat("#", w+2))
}
