package main

import (
	"fmt"
)

const (
	NODE  = 7
	EMPTY = -1
)

var (
	rear  = -1
	front = -1
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
	queue = []int{0, 0, 0, 0, 0, 0, 0}
)

func main() {
	enqueue(0)

	for {
		node := dequeue()

		if node == EMPTY {
			break
		}

		for j := 0; j < NODE; j++ {
			if graph[node][j] == 1 && visit[j] == 0 {
				enqueue(j)
			}
		}
	}
}

func enqueue(j int) {
	if rear <= NODE {
		visit[j] = 1
		fmt.Printf("%c ", j+'A')

		rear++
		queue[rear] = j
	}
}

func dequeue() int {
	front++
	if front <= rear {
		return queue[front]
	}
	return EMPTY
}
