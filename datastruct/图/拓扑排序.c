#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "graph.h"
#include "邻接表.c"

/* AOV网（Activity On Vertex NetWork，用顶点表示活动的网）：

用 DAG图（有向无环图 Directed Acyclic graph）表示一个工程。顶点表示活动，有向边
<V_i,V_j> 表示活动 V_i 必须先于活动 V_j 进行
 */

/* 拓扑排序：在图论中，由一个 DAG图
的顶点组成的序列，当且仅当满足下列条件时，称为该图的一个拓扑排序：

1.每个顶点出现且只出现一次

2.若顶点A在序列中排在顶点B的前面，则在图中不存在从顶点B到顶点A的路径
 */

/* 拓扑排序的实现：

1.从AOV网中选择一个没有前驱的顶点并输出

2.从网中删除该顶点和所有以它为起点的有向边

3.重复1和2直到当前的AOV网为空或当前网中不存在无前驱的顶点为止
 */

/* 时间复杂度
这种算法每个顶点都需要处理一次
Pop(S,i)，每条边都需要处理一次（v=p−>adjvex）。故时间复杂度为
O(|V|+|E|)，若采用邻接矩阵，则需 O(|V|^2)
 */
bool TopologicalSort(ALGraph g) {
    const int vn = g.vexnum;
    int indegree[vn];  // 每个顶点的入度，初始化过程略
    int res[vn];       // 记录拓扑排序，每个元素初始化-1
    int stack[vn];     // 存储度为 0 的顶点，也可以用队列

    for (int i = 0; i < vn; i++) {
        if (indegree[i] == 0) {
            push(stack, i);
        }
        res[i] = -1;
    }

    int count = 0;
    while (len(stack) > 0) {
        int v = pop(stack);
        res[count++] = v;

        int nbors[100];
        int nn = neighbors(v, nbors);
        for (ArcNode* p = g.vertices[v].first; p != NULL; p = p->next) {
            v = p->adjvex;
            if (--indegree[v] == 0) {
                push(stack, v);
            }
        }
    }

    // 若不等说明有回路，排序失败
    return count == vn;
}

/* 逆拓扑排序
对一个 AOV 网，如果采用下列步骤进行排序，则称之为逆拓扑排序：

1.从 AOV网中选择一个没有后继（出度为0）的顶点并输出

2.从网中删除该顶点和所有以它为终点的有向边

3.重复1和2直到当前当前的 AOV网为空

实现逆拓扑排序可以模仿拓扑排序，只不过拓扑排序中我们看的是一个结点的入度，而逆拓扑排序中我们看的是一个结点的出度，
用邻接表实现逆拓扑排序比较低效（找到指向一个顶点的边需要遍历整个邻接表），用邻接矩阵会更方便些。当然也可以用逆邻接表。
下面使用另一种方法 DFS 实现逆拓扑排序。

⚠️如果存在回路，则不存在逆拓扑排序序列，如何判断回路？如果想要访问的结点已在栈中，则有回路
*/

bool visited[100];
void DFS(graph g, int v);

void InverseTopologicalSort_DFS(graph g) {  // 对图g进行深度优先遍历
    for (int v = 0; v < g.vexnum; v++) {
        visited[v] = false;  // 初始化已访问标记数据
    }
    for (int v = 0; v < g.vexnum; ++v) {  // 本代码是从v=0开始遍历
        if (!visited[v]) {
            DFS(g, v);
        }
    }
}

void DFS(graph g, int v) {  // 从顶点v出发，深度优先遍历图g
    visited[v] = true;      // 设已访问标记
    for (int w = FirstNeighbor(g, v); w >= 0; w = NextNeighbor(g, v, w)) {
        if (!visited[w]) {  // w为u的尚未访问的邻接顶点
            DFS(g, w);
        }
    }
    print(v);  // 输出顶点
}
