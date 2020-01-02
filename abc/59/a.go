package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b, c string
	fmt.Scan(&a, &b, &c)
	s := fmt.Sprintf("%c%c%c", a[0], b[0], c[0])
	fmt.Println(strings.ToUpper(s))
}
