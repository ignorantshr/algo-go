#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "邻接表.c"

/* 拓扑排序：在图论中，由一个有向无环图（Directed Acyclic
Graph，DAG）的顶点组成的序列，当且仅当满足下列条件时，称为该图的一个拓扑排序：

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
            if (--indegree[p->adjvex] == 0) {
                push(stack, p->adjvex);
            }
        }
    }

    // 若不等说明有回路，排序失败
    return count == vn;
}