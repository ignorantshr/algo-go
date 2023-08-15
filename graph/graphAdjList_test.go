package graph

import (
	"fmt"
	"testing"
)

func TestNewGraphAdjList(t *testing.T) {
	g := NewGraphAdjList([][2]int{{5, 2}, {2, 4}, {1, 3}})
	g.AddVertex(6)
	g.AddVertex(6)
	g.AddEdge(4, 6)
	g.AddEdge(5, 6)
	g.Println()

	g.DelVertex(2)
	g.DelVertex(2)
	g.DelVertex(5)
	g.Println()
	g.DelVertex(6)
	g.DelEdge(1, 3)
	g.DelEdge(1, 3)
	g.Println()
}

func TestGraphBFS(t *testing.T) {
	g := NewGraphAdjList([][2]int{{5, 2}, {2, 4}, {1, 3}})
	g.AddVertex(6)
	g.AddVertex(6)
	g.AddEdge(4, 6)
	g.AddEdge(5, 6)
	g.Println()
	fmt.Println("-----------------")
	fmt.Println(GraphBFS(g, NewVertex(4)))
}

func TestGraphDFS(t *testing.T) {
	g := NewGraphAdjList([][2]int{{5, 2}, {2, 4}, {1, 3}, {9, 5}, {1, 2}, {3, 5}})
	g.AddVertex(6)
	g.AddVertex(6)
	g.AddEdge(4, 6)
	g.AddEdge(5, 6)
	g.Println()
	fmt.Println("-----------------")
	fmt.Println(GraphDFS(g, NewVertex(4)))
}
