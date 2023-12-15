package graph

import (
	"fmt"
)

// https://www.hello-algo.com/chapter_graph/graph/

/* 基于邻接矩阵实现的无向图类 */
type graphAdjMat struct {
	// 顶点列表，元素代表“顶点值”，索引代表“顶点索引”
	vertexes []int
	// 与 vertexes 相反，方便查询顶点的索引
	verIndex map[int]int
	// 邻接矩阵，行列索引对应“顶点索引”
	adjMat [][]int
}

func NewGraphAdjMat(edges [][3]int) *graphAdjMat {
	g := new(graphAdjMat)
	g.verIndex = make(map[int]int)
	for _, v := range edges {
		g.AddVertex(v[0])
		g.AddVertex(v[1])
		g.AddEdge(v[0], v[1], v[2])
	}
	return g
}

func (g *graphAdjMat) AddVertex(v int) {
	if has, _ := g.ExistVertex(v); has {
		return
	}

	g.vertexes = append(g.vertexes, v)
	g.verIndex[v] = len(g.vertexes) - 1

	for i := range g.adjMat {
		g.adjMat[i] = append(g.adjMat[i], 0)
	}
	newLine := make([]int, len(g.vertexes))
	g.adjMat = append(g.adjMat, newLine)
}

func (g *graphAdjMat) DelVertex(v int) {
	has, i := g.ExistVertex(v)
	if !has {
		return
	}

	length := len(g.vertexes)
	g.vertexes = append(g.vertexes[:i], g.vertexes[i+1:]...)
	for value, vi := range g.verIndex {
		if i < vi {
			g.verIndex[value]--
		}
	}
	delete(g.verIndex, v)

	for j := 0; j < length; j++ {
		if j == i {
			continue
		}
		g.adjMat[j] = append(g.adjMat[j][:i], g.adjMat[j][i+1:]...)
	}
	g.adjMat = append(g.adjMat[:i], g.adjMat[i+1:]...)
}

func (g *graphAdjMat) AddEdge(from, to, w int) {
	if from == to {
		return
	}

	ok, i := g.ExistVertex(from)
	if !ok {
		return
	}
	ok, j := g.ExistVertex(to)
	if !ok {
		return
	}
	g.adjMat[i][j] = w
	g.adjMat[j][i] = w
}

func (g *graphAdjMat) DelEdge(from, to int) {
	if from == to {
		return
	}

	ok, i := g.ExistVertex(from)
	if !ok {
		return
	}
	ok, j := g.ExistVertex(to)
	if !ok {
		return
	}
	g.adjMat[i][j] = 0
	g.adjMat[j][i] = 0
}

func (g *graphAdjMat) ExistVertex(v int) (bool, int) {
	//for i, p := range g.vertexes {
	//	if p == v {
	//		return true, i
	//	}
	//}
	//return false, -1
	val, ok := g.verIndex[v]
	return ok, val
}

func (g *graphAdjMat) Println() {
	fmt.Println("============================")
	fmt.Println("C", g.vertexes)
	for i := range g.adjMat {
		line := make([]int, 0)
		for j := range g.adjMat[i] {
			line = append(line, g.adjMat[i][j])
		}
		fmt.Println(g.vertexes[i], line)
	}
}
