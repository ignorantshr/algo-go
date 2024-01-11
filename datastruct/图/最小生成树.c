#include <assert.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "graph.h"
#include "ufset.c"

#define INF 99999

/* 对于一个带权连通无向图
G=(V,E)，生成树不同，每棵树的权（即树中所有边上的权值之和）也可能不同。设R 为G
的所有生成树的集合，若T 为R 中边的权值之和最小的生成树，则T 称为G
的最小生成树/最小代价树（Minimum−spanning−Tree,MST）
*/

/* 最小生成树的性质

最小生成树可能有多个，但边的权值总是唯一且最小的

最小生成树的边数=顶点数-1。砍掉一条则不连通，增加一条边则会出现回路

如果一个连通图本身就是一棵树，则其最小生成树就是它本身

只有连通图才有生成树，非连通图只有生成森林
 */

/* Prim（普里姆）算法

从某一个顶点开始构建生成树；每次将代价最小的新顶点纳入生成树，直到所有顶点都纳入为止，
总共需要 n-1 轮处理。

每一轮处理：循环遍历所有个结点，找到 lowCast最低的，且还没加入树的顶点。
再次循环遍历，更新还没加入的各个顶点的 lowCast 值。

每一轮处理有两次遍历循环，故时间复杂度为 O(2n) 。总时间复杂度 O(n^2)
，即O(|V|^2)，适合用于边稠密图。
*/

int Prim(graph* g) {
    int dist[MAXSIZE];
    int visited[MAXSIZE];

    for (int i = 0; i < MAXSIZE; i++) {
        dist[i] = g->edges[0][i];
        visited[i] = 0;
    }
    visited[0] = 1;  // 把一个顶点加入树

    for (int c = 1; c < MAXSIZE; c++) {  // 已经加入一个了，还需要循环 size-1 轮
        int pos = 0;
        int min = INF;
        for (int i = 0; i < MAXSIZE; i++) {
            if (!visited[i] && dist[i] < min) {  // 找一个代价最小的顶点加入树中
                pos = i;
                min = dist[i];
            }
        }

        visited[pos] = 1;
        for (int j = 0; j < MAXSIZE; j++) {
            // 这里更新的是 从刚刚纳入树中的顶点 直接到达 邻接点的最短路径
            // dist
            // 存储的是两个顶点之间的距离，而不是从某个固定的顶点到达的路径，
            // 注意与 Dijiskra 区分
            if (!visited[j] && g->edges[pos][j] < dist[j]) {
                dist[j] = g->edges[pos][j];
            }
        }
    }
    int sum = 0;
    for (int i = 0; i < MAXSIZE; i++) {
        sum += dist[i];
    }
    return sum;
}

/* Kruskal（克鲁斯卡尔）算法

每次选择一条权值最小的边，使这条边的两头连通（原本已经连通的就不选）；直到所有结点都连通（所谓连通就是两点属于同一个集合）。

需要用到并查集，刚开始把所有的点看成属于不同的集合。
初始将各条边按权值排序，接下来要把各条边检查一遍。
每一轮用并查集的知识检查该边两个顶点是否连通（是否属于同一集合），直到所有边都被检查一遍。如果发现该边两个顶点已经连通就跳过

该算法共执行 e 轮，每轮判断两个顶点是否属于同一集合，需要
O(log_2|E|)（并查集的时间复杂度）， 总的时间复杂度为
O(|E|log_2|E|)，适合用于边稀疏图
 */

typedef struct edge {
    int v1, v2;
    int weight;
} edge;

int Kruskal(graph* g) {
    edge* edges[MAXSIZE * MAXSIZE];
    int count = 0;
    // insert sort
    for (int i = 0; i < MAXSIZE; i++) {
        for (int j = 0; j < MAXSIZE; j++) {
            if (i == j) {
                continue;
            }

            if (g->edges[i][j] != INF) {
                edge* e = malloc(sizeof(edge));
                e->v1 = i;
                e->v2 = j;
                e->weight = g->edges[i][j];
                edges[count] = e;

                if (count > 0 &&
                    edges[count - 1]->weight > edges[count]->weight) {
                    edge* tmp = edges[count];
                    int k = count;
                    for (; k > 0 && edges[k - 1]->weight > tmp->weight; k--) {
                        edges[k] = edges[k - 1];
                    }
                    edges[k] = tmp;
                }
                count++;
            }
        }
    }

    int sum = 0;
    ufset* uf = initial();
    for (int i = 0; i < count; i++) {
        if (Find(uf, edges[i]->v1) != Find(uf, edges[i]->v2)) {
            sum += edges[i]->weight;
            Union(uf, edges[i]->v1, edges[i]->v2);
        }
    }

    return sum;
}

void testing() {
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

    int s1 = Prim(&g);
    printf("%d\n", s1);

    int s2 = Kruskal(&g);
    printf("%d\n", s2);

    assert(s1 == s2 && s2 == 9);
}

int main(int argc, char const* argv[]) {
    testing();
    return 0;
}
