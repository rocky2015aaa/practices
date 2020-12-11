package main

import (
	"fmt"
)

var (
	top   = -1
	graph = [][]int{
		{0, 1, 1, 0, 0, 1, 0},
		{1, 0, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0},
		{0, 0, 1, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 1, 0},
		{1, 0, 1, 1, 1, 0, 1},
		{0, 0, 0, 0, 0, 1, 0},
	}
	visit = []int{0, 0, 0, 0, 0, 0, 0}
)

func main() {
	dfs(0)
}

func dfs(v int) {
	//visit[v] = 1
	fmt.Printf("%c ", v+'A')

	for i := 0; i < len(visit); i++ {
		if graph[v][i] == 1 && visit[i] == 0 {
			visit[i] = 1
			visit[0] = 1
			dfs(i)
			visit[i] = 0
			visit[0] = 0
		}
	}
}
