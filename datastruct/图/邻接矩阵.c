#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MaxVertexNum 100  // 最大顶点数量

typedef struct mGraph {
    char vex[MaxVertexNum];
    int edge[MaxVertexNum][MaxVertexNum];
    int vexnum, arcnum;  // 图的当前顶点数和边数/弧数
};

/* 图的度

无向图的度
    第 i 个结点的度 = 第 i 行/列非零元素的个数

有向图的度
    第 i 个结点的出度：第 i 行非零元素的个数
    第 i 个结点的入度：第 i 列非零元素的个数
    第 i 个结点的度 = 第 i 行，第 i 列的非零元素个数之和

邻接矩阵法求顶点的度/出度/入度的时间复杂度都为 O(∣V∣)
空间复杂度：O(|V|^2)——只与顶点数有关，和实际的边数无关（注：O(|V|)+O(|V|^2) ）

适合存储稠密图。

无向图的邻接矩阵为对称矩阵，可以压缩存储（只存储上三角区/下三角区）

邻接矩阵法的性质 https://blog.csdn.net/weixin_51252109/article/details/131534761
设图 G 的邻接矩阵为 A（矩阵元素为  0/1 ），则 A^n[i][j] (n 个 A 相乘)等于由顶点
i 到顶点 j 的长度为 n 的路径的数目。从自己直接到自己的路径是 0。
 */

// 带权图/网
// 可用int的上限值表示“无穷”，在带权值当中如果一个值为0或无穷表示与之对应的两个顶点之间不存在边
#define MaxPowerVertexNum 100  // 顶点数目的最大值
#define INFINITY INT64_MAX     // 宏定义常量“无穷”
typedef char VertexType;       // 顶点的数据类型
typedef int EdgeType;          // 带权图中边上权值的数据类型
typedef struct {
    VertexType Vex[MaxPowerVertexNum];                    // 顶点
    EdgeType Edge[MaxPowerVertexNum][MaxPowerVertexNum];  // 边的权
    int vexnum, arcnum;  // 图的当前顶点数和弧数
} MPGraph;
