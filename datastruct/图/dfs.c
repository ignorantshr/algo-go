#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../栈和队列/队列-链式存储.c"
#include "graph.h"

// 图的深度优先遍历类似树的先根遍历，都可以用递归来实现。只不过树的先根遍历新找到的结点一定是没有访问过的

#define MAX_VERTEX_NUM 10

static bool visited[MAX_VERTEX_NUM];  // 标记访问数组

void visit(int v) {
    printf("%d ", v);
}

void DFS(graph g, int v);

/* 对图G进行深度优先遍历，对于非连通图保证结点都可被遍历到
 */
void DFSTraverse(graph g) {
    for (int i = 0; i < g.vexnum; i++) {  // 访问标记数组初始化
        visited[i] = false;
    }

    for (int i = 0; i < g.vexnum; i++) {  // 从0号顶点开始遍历
        if (!visited[i]) {  // 对每个连通分量调用一次DFS
            DFS(g, i);
        }
    }
}

void DFS(graph g, int v) {
    visit(v);
    visited[v] = true;
    for (int w = FirstNeighbor(g, v); w >= 0; w = NextNeighbor(g, v, w)) {
        if (!visited[w]) {
            DFS(g, w);
        }
    }
}

/* 空间复杂度

深度优先搜索的空间复杂度主要来自于函数的递归调用

最坏情况为单支树，递归深度为 O(|V|) 。
最好的情况为 O(1)。恰好对应广度优先遍历最坏的情况。
 */

/* 时间复杂度，同 BFS

时间复杂度=访问各结点所需时间+探索各条边所需时间

邻接矩阵存储的图：
    访问 |V| 个顶点需要 O(|V|) 的时间。查找每个顶点的邻接点 都需要  O(|V|)
的时间，总共有 |V| 个顶点，故时间复杂度= O(|V|^2)

邻接表存储的图：
    访问 |V| 个顶点需要 O(|V|) 的时间。查找每个顶点的邻接点 共需要  O(|E|)
的时间 故时间复杂度= O(|V|+|E|)
 */

/* 深度优先生成树

同一个图的邻接矩阵表示方式唯一，因此深度优先遍历序列唯一，对应深度优先生成树也唯一。

同一个图邻接表表示方法不唯一，因此深度优先遍历序列不唯一，对应深度优先生成树也不唯一

深度优先生成森林：类比广度优先生成森林
 */