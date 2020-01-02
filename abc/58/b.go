package main

import "fmt"

func main() {
	var o, e string
	fmt.Scan(&o, &e)

	ans := ""
	for i := range e {
		ans += string(o[i])
		ans += string(e[i])
	}
	if len(o) > len(e) {
		ans += string(o[len(o)-1])
	}

	fmt.Println(ans)
}
