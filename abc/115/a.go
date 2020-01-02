package main

import "fmt"

func main() {
	var day int
	fmt.Scan(&day)

	var str string
	switch day {
	case 22:
		str = "Christmas Eve Eve Eve"
	case 23:
		str = "Christmas Eve Eve"
	case 24:
		str = "Christmas Eve"
	case 25:
		str = "Christmas"
	}
	fmt.Println(str)
}
