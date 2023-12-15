package graph

import (
	"fmt"
	"sort"
)

/*
最小生成树算法

	Kruskal（克鲁斯卡尔）算法

每次选择一条权值最小的边，使这条边的两头连通（原本已经连通的就不选）；直到所有结点都连通
时间复杂度为 O(|E|log_2|E|) ，适合用于边稀疏图
*/
func Kruskal(g graphAdjMat) {
	// 按边排序，并查集
	edges := make([]edge, 0)

	// 下面是无向图的解法，搜索矩阵右上三角区
	for i := 0; i < len(g.adjMat); i++ {
		for j := i; j < len(g.adjMat[0]); j++ {
			if g.adjMat[i][j] > 0 {
				edges = append(edges, edge{
					i, j, g.adjMat[i][j],
				})
			}
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	uf := NewUF(len(g.vertexes))
	for _, edge := range edges {
		if uf.Find(edge.v1) != uf.Find(edge.v2) {
			addTree(edge.v1, edge.v2, edge.weight)
			uf.Union(edge.v1, edge.v2)
		}
	}
}

type edge struct {
	v1, v2 int
	weight int
}

func addTree(v1, v2, w int) {
	fmt.Printf("[%d--%d--%d]\n", v1, w, v2)
}
