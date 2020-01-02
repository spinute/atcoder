package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	ans := ""
	for i, cnt := 0, 1; i < len(s); i++ {
		if i+1 == len(s) || s[i] != s[i+1] {
			ans += string(s[i]) + strconv.Itoa(cnt)
			cnt = 1
		} else {
			cnt++
		}
	}

	fmt.Println(ans)
}
