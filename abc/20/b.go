package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	ret, _ := strconv.Atoi(a + b)
	fmt.Println(ret * 2)
}
