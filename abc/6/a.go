package main

import "fmt"

var p = fmt.Println

func main() {
	var a int
	fmt.Scan(&a)
	if a%3 == 0 {
		p("YES")
	} else {
		p("NO")
	}
}
