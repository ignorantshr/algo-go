package graph

import (
	"testing"
)

func TestNewGraphAdjMat(t *testing.T) {
	g := NewGraphAdjMat([][2]int{{5, 2}, {2, 4}, {1, 3}})
	g.AddVertex(6)
	g.AddVertex(6)
	g.AddEdge(4, 6)
	g.AddEdge(5, 6)
	g.Println()

	g.DelVertex(2)
	g.DelVertex(2)
	g.DelVertex(5)
	g.DelVertex(6)
	g.DelEdge(1, 3)
	g.DelEdge(1, 3)
	g.Println()
}
