package graph

import (
	"testing"
)

func TestNewGraphAdjMat(t *testing.T) {
	g := NewGraphAdjMat([][3]int{{5, 2, 1}, {2, 4, 1}, {1, 3, 1}})
	g.AddVertex(6)
	g.AddVertex(6)
	g.AddEdge(4, 6, 1)
	g.AddEdge(5, 6, 1)
	g.Println()

	g.DelVertex(2)
	g.DelVertex(2)
	g.DelVertex(5)
	g.DelVertex(6)
	g.DelEdge(1, 3)
	g.DelEdge(1, 3)
	g.Println()
}
