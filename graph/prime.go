package graph

import "math"

/*
最小生成树算法

	Prim（普里姆）算法

从某一个顶点开始构建生成树；每次将代价最小的新顶点纳入生成树，直到所有顶点都纳入为止
时间复杂度为 O(|V^2|) ，适合用于边稠密图
*/
func Prime(g graphAdjMat) int {
	// 下面是无向图的解法
	size := len(g.vertexes)
	cost := make([]int, size)
	visited := make([]int, size)
	added := 1
	lasti := 0

	visited[0] = 1
	// 第一轮遍历
	for i := 0; i < size; i++ {
		if v := g.adjMat[0][i]; visited[i] == 0 && v != 0 {
			cost[i] = v
		} else {
			cost[i] = math.MaxInt
		}
	}

	for added < size {
		nidx := 0
		// 找到还未加入到树中的代价最小的顶点
		for i := 0; i < size; i++ {
			if visited[i] == 0 && cost[i] < cost[nidx] {
				nidx = i
			}
		}
		visited[nidx] = 1 // 访问标记
		added++
		addTree(lasti, nidx, g.adjMat[lasti][nidx])
		lasti = nidx

		// 遍历顶点的边
		for i := 0; i < size; i++ {
			// 如果从自己这里中转比原来的路径更短，那么更新
			if v := g.adjMat[nidx][i]; visited[i] == 0 && g.adjMat[nidx][i] != 0 && v < cost[i] {
				cost[i] = v
			}
		}
	}

	sum := 0
	for i := 0; i < size; i++ {
		sum += cost[i]
	}
	return sum
}
