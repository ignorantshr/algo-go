#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "graph.h"

#define INF 99999
void distAndPath(int s, int distance[], int path[]);
void printSolution(int dist[MAXSIZE][MAXSIZE], int path[MAXSIZE][MAXSIZE]);
void findPath(int i, int j, int path[MAXSIZE][MAXSIZE]);

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

// 邻接矩阵实现
void BFS_MIN_Distance(graph* g, int s) {
    int dist[MAXSIZE];
    int path[MAXSIZE];
    bool visited[MAXSIZE];

    int queque[MAXSIZE];
    int front = 0, rear = -1;  // 模拟队列的头尾

    for (int i = 0; i < MAXSIZE; i++) {
        dist[i] = INF;
        path[i] = INF;
        visited[i] = false;
    }

    queque[++rear] = s;
    visited[s] = true;
    dist[s] = 0;

    while (front <= rear) {
        int v = queque[front];
        for (int j = 0; j < MAXSIZE; j++) {
            if (!visited[j] && g->edges[v][j] != INF) {
                visited[j] = true;
                dist[j] = dist[v] + 1;
                path[j] = v;
                queque[++rear] = j;
            }
        }
        front++;
    }

    distAndPath(s, dist, path);
}

// 邻接表实现，只是参考代码
/* void BFS_MIN_Distance(graph g, int u) {
    // d[i]表示从u到i结点的最短路径
    int d[MAXSIZE];
    int path[MAXSIZE];
    bool visited[MAXSIZE];
    LinkQueue q;

    for (int i = 0; i < MAXSIZE; i++) {
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
 */

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

// 邻接矩阵有向图 具体实现
void Dijkstra(graph* g, int s) {
    bool final[MAXSIZE];    // 顶点是否确定了
    int distance[MAXSIZE];  // 路径长度
    int path[MAXSIZE];      // 路径
    for (int i = 0; i < g->vexnum; i++) {
        final[i] = false;
        distance[i] = INF;
        path[i] = INF;
    }

    distance[s] = 0;  // 起点

    // 每回合在 distance 中确定一个 未确定的 最短路径的 结点 next
    // 遍历它的邻接点，如果经过它中转之后比从 s
    // 经过其它路径到达更短，则更新其路径
    for (int c = 0; c < g->vexnum; c++) {
        int next = -1;
        int mincost = INF;
        // 找到一个最短路径
        for (int i = 0; i < g->vexnum; i++) {
            if (!final[i] && distance[i] <= mincost) {
                next = i;
                mincost = distance[i];
            }
        }

        final[next] = true;
        for (int j = 0; j < g->vexnum; j++) {
            if (!final[j] && g->edges[next][j] != INF &&
                distance[next] + g->edges[next][j] < distance[j]) {
                //
                distance[j] = distance[next] + g->edges[next][j];
                path[j] = next;
            }
        }
    }

    distAndPath(s, distance, path);
}

void distAndPath(int s, int distance[], int path[]) {
    printf("-------distance--------\n");
    for (int i = 0; i < MAXSIZE; i++) {
        printf("%d ", distance[i]);
    }
    printf("\n");

    printf("-------path--------\n");
    for (int i = 0; i < MAXSIZE; i++) {
        printf("%d ", path[i]);
    }
    printf("\n");

    // 若想求某个具体路径则需要反推
    for (int c = 0; c < MAXSIZE; c++) {
        if (c == s) {
            continue;
        }

        int end = c;
        int pos = MAXSIZE;
        int p[MAXSIZE];
        while (path[end] != INF) {
            p[--pos] = end;
            end = path[end];
        }
        if (pos < MAXSIZE) {
            p[--pos] = s;
        }

        while (pos < MAXSIZE) {
            printf("%d->", p[pos++]);
        }
        printf("\n");
    }
}

/* Floyd算法

动态规划思想。假设我们正在寻找从点 i 到点 j 的最短路径，
如果我们知道从 i 到一个中间点 k 以及从 k 到 j 的最短路径，
那么 i 到 j 的最短路径可能就是这两段路径的叠加，
前提是这段路径比直接从 i 到 j 的路径短。

Floyd算法可以用于负权值带权图，但不能解决带有“负权回路”的图（有负权值的边组成回路），这种图可能没有最短路径

时间复杂度： O(|V|^3)

空间复杂度：O(|V|^2)
 */

void Floyd_MIN_Distance(graph* g, int u) {
    int dist[MAXSIZE][MAXSIZE];  // 从 i 到 j 的距离
    int path[MAXSIZE][MAXSIZE];
    //......准备工作，根据图的信息初始化矩阵A和path
    for (int i = 0; i < MAXSIZE; i++) {
        for (int j = 0; j < MAXSIZE; j++) {
            dist[i][j] = g->edges[i][j];
            if (g->edges[i][j] == INF) {
                path[i][j] = INF;
            } else {
                path[i][j] = i;
            }
        }
    }

    for (int k = 0; k < MAXSIZE; k++) {  // 考虑以V_k作为中转点
        for (int i = 0; i < MAXSIZE; i++) {  // 遍历整个矩阵，i为行号，j为列号
            for (int j = 0; j < MAXSIZE; j++) {
                if (dist[i][k] < INF && dist[i][k] < INF &&
                    dist[i][k] + dist[k][j] < dist[i][j]) {
                    // 如果经过 i->k->j 的路径更短，则更新从 i->j 的距离
                    dist[i][j] = dist[i][k] + dist[k][j];
                    path[i][j] = k;
                }
            }
        }
    }

    printSolution(dist, path);
    findPath(0, 3, path);
}

void printSolution(int dist[MAXSIZE][MAXSIZE], int path[MAXSIZE][MAXSIZE]) {
    printf(
        "Following matrix shows the shortest distances between every pair of "
        "vertices \n");
    printf("-------distance--------\n");
    for (int i = 0; i < MAXSIZE; i++) {
        for (int j = 0; j < MAXSIZE; j++) {
            if (dist[i][j] == INF)
                printf("%7s", "INF");
            else
                printf("%7d", dist[i][j]);
        }
        printf("\n");
    }

    printf("-------path--------\n");
    for (int i = 0; i < MAXSIZE; i++) {
        for (int j = 0; j < MAXSIZE; j++) {
            if (path[i][j] == INF)
                printf("%7s", "INF");
            else
                printf("%7d", path[i][j]);
        }
        printf("\n");
    }
}  // 0-1 (0-0 0-1) 1-2 (0-1 1-2)

void findPath(int i, int j, int path[MAXSIZE][MAXSIZE]) {
    int route[MAXSIZE];
    int pos = MAXSIZE;
    for (int cur = j; cur != i; cur = path[i][cur]) {
        route[--pos] = cur;
    }
    if (pos < MAXSIZE) {
        route[--pos] = i;
    }

    while (pos < MAXSIZE) {
        printf("%d->", route[pos++]);
    }
    printf("\n");
}

void testingBFS() {
    int edges[MAXSIZE][MAXSIZE] = {
        {0, 1, INF, 1}, {INF, 0, 1, INF}, {INF, INF, 0, 1}, {INF, INF, INF, 0}};
    graph g = {.vexnum = MAXSIZE};
    for (int i = 0; i < MAXSIZE; i++) {
        for (int j = 0; j < MAXSIZE; j++) {
            g.edges[i][j] = edges[i][j];
        }
    }

    BFS_MIN_Distance(&g, 0);
}

void testingDijkstra() {
    int edges[MAXSIZE][MAXSIZE] = {{0, 5, INF, 10},
                                   {INF, 0, 3, INF},
                                   {INF, INF, 0, 1},
                                   {INF, INF, INF, 0}};
    graph g = {.vexnum = MAXSIZE};
    for (int i = 0; i < MAXSIZE; i++) {
        for (int j = 0; j < MAXSIZE; j++) {
            g.edges[i][j] = edges[i][j];
        }
    }

    Dijkstra(&g, 0);
}

void testingFloyd() {
    int edges[MAXSIZE][MAXSIZE] = {{0, 5, INF, 10},
                                   {INF, 0, 3, INF},
                                   {INF, INF, 0, 1},
                                   {INF, INF, INF, 0}};
    graph g = {.vexnum = MAXSIZE};
    for (int i = 0; i < MAXSIZE; i++) {
        for (int j = 0; j < MAXSIZE; j++) {
            g.edges[i][j] = edges[i][j];
        }
    }

    Floyd_MIN_Distance(&g, 0);
}

int main(int argc, char const* argv[]) {
    // testingBFS();
    // testingDijkstra();
    testingFloyd();
    return 0;
}
