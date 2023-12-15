package graph

import (
	"fmt"
)

// 并查集 union find set

type uf struct {
	parent []int
}

func NewUF(size int) *uf {
	u := &uf{make([]int, size)}
	for i := 0; i < size; i++ {
		u.parent[i] = -1
	}
	return u
}

// 查找 x 所属的集合（根节点）
func (u *uf) Find(i int) int {
	r := i
	for u.parent[r] != -1 {
		r = u.parent[r]
	}

	for i != r {
		i, u.parent[i] = u.parent[i], r
	}

	return r
}

func (u *uf) Union(i1, i2 int) {
	r1 := u.Find(i1)
	r2 := u.Find(i2)

	if r1 == r2 {
		return
	}
	u.parent[r1] = r2
}

func (u *uf) print() {
	fmt.Println("----------")
	for i := 0; i < len(u.parent); i++ {
		fmt.Printf("%d ", u.parent[i])
	}
	fmt.Println()
	fmt.Println("----------")
}

func (u *uf) findPrint() {
	fmt.Println("-----find-----")
	for i := 0; i < len(u.parent); i++ {
		fmt.Printf("%d ", u.Find(i))
	}
	fmt.Println()
	fmt.Println("-----find-----")
}
