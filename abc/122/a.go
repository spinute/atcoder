package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)
	switch n {
	case "A":
		fmt.Println("T")
	case "T":
		fmt.Println("A")
	case "C":
		fmt.Println("G")
	case "G":
		fmt.Println("C")
	}
}
