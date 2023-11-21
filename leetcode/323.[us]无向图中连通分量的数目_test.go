/*
- @lc app=leetcode.cn id=323 lang=golang

给你输入一个包含 n 个节点的图，用一个整数 n 和一个数组 edges 表示，其中 edges[i] = [ai, bi] 表示图中节点 ai 和 bi 之间有一条边。
请你计算这幅图的连通分量个数。
*/
package leetcode

import (
	"fmt"
	"testing"
)

// @lc code=start
func countComponents(n int, edges [][]int) int {
	uf := InitUF323(n)
	for i := 0; i < len(edges); i++ {
		uf.Union(edges[i][0], edges[i][1])
	}
	return uf.Count()
}

// 并查集
type UF323 struct {
	parent []int // 父节点索引
	count  int   // 连通分量
}

func InitUF323(n int) *UF323 {
	u := &UF323{
		parent: make([]int, n),
		count:  n,
	}
	for i := 0; i < n; i++ {
		u.parent[i] = -1
	}
	return u
}

func (u *UF323) Find(x int) int {
	r := x
	for u.parent[r] > 0 {
		r = u.parent[r]
	}

	for u.parent[x] > 0 {
		// tmp := u.parent[x]
		x, u.parent[x] = u.parent[x], r
		// x = tmp
	}

	return x
}

func (u *UF323) Union(x1, x2 int) {
	r1 := u.Find(x1)
	r2 := u.Find(x2)

	if r1 == r2 {
		return
	}

	u.count--
	u.parent[r1] = r2
}

func (u UF323) Count() int {
	return u.count
}

func (u UF323) Connected(x1, x2 int) bool {
	return u.Find(x1) == u.Find(x2)
}

// @lc code=end

func Test_countComponents(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{
			10,
			[][]int{},
		}, 10},
		{"", args{
			10,
			[][]int{{1, 2}},
		}, 9},
		{"", args{
			10,
			[][]int{{1, 2}, {2, 0}, {9, 8}},
		}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countComponents(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("countComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUF(t *testing.T) {
	n := 10
	uf := InitUF323(n)
	uf.Union(0, 1)
	printUF(*uf)
	uf.Union(1, 2)
	printUF(*uf)
	fmt.Printf("root: %d\n", uf.Find(0))
	printUF(*uf)
}

func printUF(uf UF323) {
	fmt.Printf("---------------%d---------------\n", uf.Count())
	for i := 0; i < len(uf.parent); i++ {
		fmt.Printf("%d ", uf.parent[i])
	}
	fmt.Println()
}
