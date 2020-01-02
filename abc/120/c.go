package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	count := 0
	ans := 0
	for _, r := range s {
		if r == '0' {
			if count > 0 {
				ans += 2
			}
			count--
		} else {
			if count < 0 {
				ans += 2
			}
			count++
		}
	}
	fmt.Println(ans)
}
