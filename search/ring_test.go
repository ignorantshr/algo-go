package search

import (
	"container/ring"
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	logring := func(r *ring.Ring) {
		fmt.Println("----------")
		r.Do(func(a any) { fmt.Print(a, "->") })
		fmt.Println("\n----------")
	}

	r := ring.New(3)

	n := r
	n.Value = 1
	n = n.Next()
	n.Value = 2
	n = n.Next()
	n.Value = 3
	// logring(r)

	r2 := ring.New(2)
	n2 := r2
	n2.Value = 7
	n2 = n2.Next()
	n2.Value = 8
	// logring(r2)

	r = n.Link(r2)
	// logring(r)

	r3 := r.Move(2)
	logring(r3)

	r.Unlink(1)
	logring(r)
}
