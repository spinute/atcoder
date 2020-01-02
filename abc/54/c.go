package main

import "fmt"

func all(l []bool) bool {
	for _, v := range l {
		if !v {
			return false
		}
	}
	return true
}

func dfs(g [][]int, node int, visited []bool) int {
	if all(visited) {
		return 1
	}

	sum := 0
	for _, neigh := range g[node] {
		if !visited[neigh] {
			visited[neigh] = true
			sum += dfs(g, neigh, visited)
			visited[neigh] = false
		}
	}
	return sum
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	as := make([]int, m)
	bs := make([]int, m)
	for i := range as {
		fmt.Scan(&as[i])
		fmt.Scan(&bs[i])
	}

	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, 0)
	}
	for i, a := range as {
		b := bs[i]
		g[a-1] = append(g[a-1], b-1)
		g[b-1] = append(g[b-1], a-1)
	}

	visited := make([]bool, n)
	visited[0] = true
	fmt.Println(dfs(g, 0, visited))
}
