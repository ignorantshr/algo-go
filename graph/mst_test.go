package graph

import (
	"fmt"
	"testing"
)

func TestPrime(t *testing.T) {
	g := NewGraphAdjMat([][3]int{{0, 1, 1}, {0, 2, 8}, {1, 2, 3}, {2, 3, 1}})
	g.Println()

	fmt.Println("------Kruskal------")
	Kruskal(*g)
	fmt.Println("------Prime------")
	Prime(*g)
}

func TestUF(t *testing.T) {
	u := NewUF(5)
	u.Union(1, 2)
	u.print()
	u.findPrint()
	u.Union(2, 4)
	u.print()
	u.findPrint()
}
