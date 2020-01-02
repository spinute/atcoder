package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ss := make([]string, n)
	for i := range ss {
		fmt.Scan(&ss[i])
	}
	counter := make(map[string]int)
	for _, s := range ss {
		counter[s]++
	}

	if len(counter) == 3 {
		fmt.Println("Three")
	} else {
		fmt.Println("Four")
	}
}
