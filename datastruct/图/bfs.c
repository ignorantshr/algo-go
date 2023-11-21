#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../栈和队列/队列-链式存储.c"
#include "graph.h"

#define MAX_VERTEX_NUM 10

static bool visited[MAX_VERTEX_NUM];  // 标记访问数组
static LinkQueue q;

// 需要借助图的基本操作中的 FirstNeighbor(G,x) NextNeighbor(G,x,y)
// 需要设置访问标记数组 visited，初始时元素都置为 false表示结点未访问

void visit(int v) {
    printf("%d ", v);
}

void BFS(graph g, int v);

/* 对图G进行广度优先遍历，对于非连通图保证结点都可被遍历到
对于无向图，调用BFS函数的次数=连通分支数

对于有向图，即使只有一个连通分支，从不同的顶点出发也不一定都能一次性遍历完所有顶点，不同顶点调用BFS函数的次数不一定相等
*/

void BFSTraverse(graph g) {
    InitQueue1(&q);

    for (int i = 0; i < g.vexnum; i++) {  // 访问标记数组初始化
        visited[i] = false;
    }

    for (int i = 0; i < g.vexnum; i++) {  // 从0号顶点开始遍历
        if (!visited[i]) {  // 对每个连通分量调用一次BFS
            BFS(g, i);
        }
    }
}

// 广度优先遍历
void BFS(graph g, int v) {
    visit(v);
    visited[v] = true;
    enQueue2(&q, v);  // 从顶点v出发，广度优先遍历图G
    while (!empty2(q)) {
        deQueue2(&q, &v);
        for (int w = FirstNeighbor(g, v); w >= 0; w = NextNeighbor(g, v, w)) {
            if (!visited[w]) {
                visit(w);
                visited[w] = true;
                enQueue2(&q, w);
            }
        }
    }
}

/* 空间复杂度

BFS算法的空间复杂度主要来自辅助队列

最坏情况下一个结点连接其余所有结点，辅助队列大小为 O(|V|)
 */

/* 时间复杂度

时间复杂度主要来自访问各个顶点和探索各条边

邻接矩阵存储的图：
    访问 |V| 个顶点需要 O(|V|) 的时间。查找每个顶点的邻接点 都需要  O(|V|)
的时间，总共有 |V| 个顶点，故时间复杂度= O(|V|^2)

邻接表存储的图：
    访问 |V| 个顶点需要 O(|V|) 的时间。查找每个顶点的邻接点 共需要  O(|E|)
的时间 故时间复杂度= O(|V|+|E|)
 */

/* 广度优先生成树

根据广度优先遍历过程中访问结点的先后顺序依次将未访问（未入队）的结点插入构造广度优先生成树。

广度优先生成树由广度优先遍历的过程确定。
邻接矩阵法生成的广度优先生成树唯一。
由于邻接表表示方法不唯一，因此基于邻接表的广度优先生成树也不唯一。

对非连通图的广度优先遍历，可得广度优先生成森林
 */
