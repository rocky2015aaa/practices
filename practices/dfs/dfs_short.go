package main

import (
	"fmt"
)

var (
	graph = [][]int{
		{0, 1, 1, 0, 0, 1, 0},
		{1, 0, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0},
		{0, 0, 1, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 1, 0},
		{1, 0, 1, 1, 1, 0, 1},
		{0, 0, 0, 0, 0, 1, 0},
	}
	min = 0
)

func main() {
	min = 7 * 7
	dfs(0, 0, 1)
	fmt.Println(min)
}

func dfs(x, y, l int) {
	if x == 6 && y == 6 {
		if min > l {
			min = l
			return
		}
	}
	graph[y][x] = 0
	if y > 0 && graph[y-1][x] != 0 {
		dfs(x, y-1, l+1)
	}
	if y < 6 && graph[y+1][x] != 0 {
		dfs(x, y+1, l+1)
	}
	if x > 0 && graph[y][x-1] != 0 {
		dfs(x-1, y, l+1)
	}
	if x < 6 && graph[y][x+1] != 0 {
		dfs(x+1, y, l+1)
	}

	graph[y][x] = 1
}
