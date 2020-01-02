package main

import "fmt"

func main() {
	var ss [3]string

	for i := range ss {
		fmt.Scan(&ss[i])
	}

	for i := range ss {
		fmt.Printf("%c", ss[i][i])
	}
	fmt.Println("")
}
