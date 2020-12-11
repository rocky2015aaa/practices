package main

import (
	"fmt"
)

const NODE = 7

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
	stack = []int{0, 0, 0, 0, 0, 0, 0}
)

func main() {
	push(0)

	for {
		for j := 0; j < NODE; j++ {
			if graph[stack[top]][j] == 1 && visit[j] == 0 {
				push(j)
				j = 0
			}
		}
		pop()
		if top < 0 {
			fmt.Println()
			break
		}
	}
}

func push(j int) {
	if top <= NODE {
		visit[j] = 1
		fmt.Printf("%c ", j+'A')

		top++
		stack[top] = j
	}
}

func pop() {
	if top >= 0 {
		top--
	}
}
