package main

import "fmt"

func main() {
	var s string
	var k int
	fmt.Scan(&s, &k)

	counter := make(map[string]int)
	for i := 0; i+k <= len(s); i++ {
		counter[s[i:i+k]]++
	}
	fmt.Println(len(counter))
}
