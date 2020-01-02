package main

import "fmt"

func score(v int) int {
	if v == 1 {
		return 14
	}
	return v
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if sa, sb := score(a), score(b); sa > sb {
		fmt.Println("Alice")
	} else if sa == sb {
		fmt.Println("Draw")
	} else {
		fmt.Println("Bob")
	}
}
