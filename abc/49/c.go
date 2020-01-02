package main

import "fmt"

func dfs(s string) bool {
	if s == "" {
		return true
	}
	for _, t := range []string{"dream", "dreamer", "erase", "eraser"} {
		if len(s) >= len(t) && s[0:len(t)] == t && dfs(s[len(t):]) {
			return true
		}
	}
	return false
}

func main() {
	var s string
	fmt.Scan(&s)

	if dfs(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
