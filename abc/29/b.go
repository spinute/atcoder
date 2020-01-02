package main

import "fmt"

func main() {
	ss := make([]string, 12)
	for i := range ss {
		fmt.Scan(&ss[i])
	}

	cnt := 0
	for _, s := range ss {
		for _, r := range s {
			if r == 'r' {
				cnt++
				break
			}
		}
	}
	fmt.Println(cnt)
}
