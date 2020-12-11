package main

import (
	"fmt"
)

type Graph struct {
	NumberOfNodes int
	Links         [][]int
	Paths         [][]int // all possible path list
	ShortestPath  []int
}

func (g *Graph) SetLink(sNode, dNode int) {
	g.Links[sNode] = append(g.Links[sNode], dNode) // index is source node to link, value is linked nodes list
}

func (g *Graph) SetAllPaths(source, dest int) {
	visited := make([]bool, g.NumberOfNodes, g.NumberOfNodes)
	path := make([]int, g.NumberOfNodes, g.NumberOfNodes)
	var pathIndex int

	g.setAllPathsHelper(source, dest, visited, path, &pathIndex)
}

/*
	- setAllPathsHelper
	* parameters: source node, dest node, visited list, path, path index

	With recursion for DFS, the logic to find all possible paths
*/
func (g *Graph) setAllPathsHelper(source, dest int, visited []bool, path []int, pathIndex *int) {
	visited[source] = true
	path[*pathIndex] = source
	*pathIndex++

	if source == dest {
		newPath := []int{}
		for i := 0; i < *pathIndex; i++ {
			newPath = append(newPath, path[i])
		}
		g.Paths = append(g.Paths, newPath)
	} else { // if current node is not destination then run the recursion with all linked nodes to current node
		for _, linkedNode := range g.Links[source] {
			if !visited[linkedNode] {
				g.setAllPathsHelper(linkedNode, dest, visited, path, pathIndex)
			}
		}
	}

	// delete current node data like path index and mark for visited list
	*pathIndex--
	visited[source] = false
}

func (g *Graph) PrintAllPaths() {
	for _, path := range g.Paths {
		for _, node := range path {
			fmt.Printf("%c ", node+'A')
		}
		fmt.Println()
	}
}

func (g *Graph) SetShortestPath() {
	g.ShortestPath = g.Paths[0]
	for i := 1; i < len(g.Paths); i++ {
		if len(g.ShortestPath) > len(g.Paths[i]) {
			g.ShortestPath = g.Paths[i]
		}
	}
}

func (g *Graph) PrintShortestPath() {
	for _, node := range g.ShortestPath {
		fmt.Printf("%c ", node+'A')
	}
	fmt.Println()
}

func main() {
	g := Graph{7, make([][]int, 7, 7), [][]int{}, []int{}}

	g.SetLink(0, 1)
	g.SetLink(0, 2)
	g.SetLink(0, 5)
	g.SetLink(1, 2)
	g.SetLink(1, 4)
	g.SetLink(2, 3)
	g.SetLink(2, 5)
	g.SetLink(3, 5)
	g.SetLink(4, 5)
	g.SetLink(5, 6)

	g.SetLink(1, 0)
	g.SetLink(2, 0)
	g.SetLink(5, 0)
	g.SetLink(2, 1)
	g.SetLink(4, 1)
	g.SetLink(3, 2)
	g.SetLink(5, 2)
	g.SetLink(5, 3)
	g.SetLink(5, 4)
	g.SetLink(6, 5)

	g.SetAllPaths(0, 6)
	g.PrintAllPaths()
	fmt.Println()
	g.SetShortestPath()
	g.PrintShortestPath()
	fmt.Println()
	fmt.Println(len(g.ShortestPath))
}
