package main

import "fmt"

var p = fmt.Println

func main() {
	var a int
	fmt.Scan(&a)
	switch {
	case a < 60:
		p("Bad")
	case a < 90:
		p("Good")
	case a < 100:
		p("Great")
	default:
		p("Perfect")
	}
}
