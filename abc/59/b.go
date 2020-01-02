package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	ans := ""
	if la, lb := len(a), len(b); la < lb {
		ans = "LESS"
	} else if la > lb {
		ans = "GREATER"
	} else {
		for i := 0; i < la; i++ {
			if a[i] == b[i] {
				continue
			}

			if a[i] > b[i] {
				ans = "GREATER"
			} else {
				ans = "LESS"
			}
			break
		}
	}
	if ans == "" {
		ans == "EQUAL"
	}
	fmt.Println(ans)
}
