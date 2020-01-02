package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	as := make([]int, n)
	for i := range as {
		fmt.Scan(&as[i])
	}
	cf := 0
	cn := 0
	ct := 0
	for _, a := range as {
		if a%4 == 0 {
			cf++
		}
		if a%2 == 0 {
			ct++
		} else {
			cn++
		}
	}
	if ct == cf {
		if cf >= cn-1 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	} else {
		if cf >= cn {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
