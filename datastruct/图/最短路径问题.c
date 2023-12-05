#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../栈和队列/队列-链式存储.c"
#include "graph.h"

/* 从 V_0 开始到其他各个顶点的最短路径。

https://blog.csdn.net/zimuzi2019/article/details/126283227

BFS算法求单源最短路径只适用于无权图，或所有边的权值都相同的图。
Dijkstra 算法和 Floyd
算法则对带权图和无权图都适用。Dijkstra算法不适用于有带负权值的带权图

其中BFS算法和Dijkstra算法用于解决单源最短路径问题，Floyd算法用于求各顶点间的最短路径
 */

/* BFS

在visit一个顶点时，修改其最短路径长度d[]并在path[]记录前驱结点。
d[]数组反映起点到目标结点的最短长度，path[]数组可以反映最短路径的走法
 */
void BFS_MIN_Distance(graph g, int u) {
    // d[i]表示从u到i结点的最短路径
    const int n = g.vexnum;
    int d[n];
    int path[n];
    bool visited[n];
    LinkQueue q;

    for (int i = 0; i < n; i++) {
        d[i] = INT64_MAX;  // 初始化路径长度
        path[i] = -1;      // 最短路径从哪个顶点过来
    }

    d[u] = 0;
    visited[u] = true;
    EnQueue(q, u);
    while (!IsEmpty(q)) {  // BFS算法主过程
        DeQueue(q, u);     // 队头元素u出队
        for (int w = FirstNeighbor(g, u); w >= 0; w = NextNeighbor(g, u, w)) {
            if (!visited[w]) {      // w为u的尚未访问的邻接顶点
                d[w] = d[u] + 1;    // 路径长度加1
                path[w] = u;        // 最短路径应从u到w
                visited[w] = true;  // 设已访问标记
                EnQueue(q, w);      // 顶点w进队
            }                       // if
        }
    }  // while
}

/* Dijkstra 算法
算法思想：
使用一个 d 数组记录距起点的最短距离，
使用一个 p 数组记录从起点到某个顶点的路径的前驱结点
每轮循环从 d 中找一个最短的还没被确定的顶点 i，
然后遍历 i 所有的邻接点，更新它们的最短路径
    若 d[i] + arc[i][j] < d[j]，说明经过 i 到达 j 比直接到达 j
要短，那么就更新最短路径信息

虽然适用于有向图，但是可以把无向图的无向边可以看成双向的两条有向边。

从 dist数组可知起点到各个结点的最短路径长度，而检查
path数组就能得到该最短路径的具体信息

时间复杂度：

O(n^2) 即 O(|V|^2)。经过 n−1轮处理，每次处理时间复杂度为 O(n)+O(n)
 */

// 仅为算法思想体现，不是完整实现
void Dijkstra_MIN_Distance(graph g, int u, int arcs[10][10]) {
    const int n = g.vexnum;
    int dist[n];
    int path[n];
    bool final[n];

    // 起点
    final[u] = true;
    dist[u] = 0;
    path[u] = -1;

    // 其余顶点
    for (int k = 0; k < n; k++) {
        if (k == u) {
            continue;
        }
        final[k] = false;
        dist[k] = arcs[u][k];  // 从 u 到 k 的弧的权值
        path[k] = (arcs[u][k] == INT64_MAX) ? -1 : 0;
    }

    // 循环遍历所有顶点，找到还没确定最短路径，且 dist最小的顶点
    int i;
    for (i = 0; i < n; i++) {
    }

    final[i] = true;
    // 检查所有邻接自 v_i 的顶点
    int neighbors[100];
    int nn = Neighbors(g, i, neighbors);
    /*
    n1---1---n2---5---n3
    n1---9---n3

    n1 = k  起始顶点 u
    n2 = i  距 u 最小的顶点
    n3 = j  i 的邻接点
     */
    for (int m = 0; m < nn; m++) {
        int j = neighbors[m];
        if (final[j] == false && dist[i] + arcs[i][j] < dist[j]) {
            dist[j] = dist[i] + arcs[i][j];
            path[j] = i;
        }
    }
}

/* Floyd算法

Floyd算法可以用于负权值带权图，但不能解决带有“负权回路”的图（有负权值的边组成回路），这种图可能没有最短路径

时间复杂度： O(|V|^3)

空间复杂度：O(|V|^2)
 */

void Floyd_MIN_Distance(graph g, int u) {
    const int n = g.vexnum;
    int A[n][n];
    int path[n][n];
    //......准备工作，根据图的信息初始化矩阵A和path
    for (int k = 0; k < n; k++) {  // 考虑以V_k作为中转点
        for (int i = 0; i < n; i++) {  // 遍历整个矩阵，i为行号，j为列号
            for (int j = 0; j < n; j++) {
                if (A[i][j] > A[i][k] + A[k][j]) {  // 以V_k为中转点的路径更短
                    A[i][j] = A[i][k] + A[k][j];  // 更新最短路径长度
                    path[i][j] = k;               // 中转点
                }
            }
        }
    }
}