package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a/10 == 9 || a%10 == 9 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
