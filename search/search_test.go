package main

import (
	"reflect"
	"testing"
)

func TestSetShortestPath(t *testing.T) {
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
	g.SetShortestPath()
	result := reflect.DeepEqual(g.ShortestPath, []int{0, 5, 6})
	if !result {
		t.Error("expected", true, "got", result)
	}
}
