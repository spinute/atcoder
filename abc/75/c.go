package main

import "fmt"

func different(i, j, node, neigh int) bool {
	return !(i == node && j == neigh) && !(i == neigh && j == node)
}

func dfs(g [][]int, i, j, node int, visited []bool) {
	for _, neigh := range g[node] {
		if !visited[neigh] && different(i, j, node, neigh) {
			visited[neigh] = true
			dfs(g, i, j, neigh, visited)
		}
	}
}

func is_bridge(g [][]int, i, j int) bool {
	visited := make([]bool, len(g))
	visited[0] = true
	dfs(g, i, j, 0, visited)
	for _, b := range visited {
		if !b {
			return true
		}
	}
	return false
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, 0)
	}
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		g[a-1] = append(g[a-1], b-1)
		g[b-1] = append(g[b-1], a-1)
	}

	cnt := 0
	for i, l := range g {
		for _, j := range l {
			if i < j && is_bridge(g, i, j) {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
