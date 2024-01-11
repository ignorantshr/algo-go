#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MaxVertexNum 5
typedef char VertexType;  // 顶点的数据类型

// 边/弧
typedef struct ArcNode {
    int adjvex;            // 顶点编号（vertices 的索引）
    struct ArcNode* next;  // 指向下一条弧的指针
    int power;             // 边权值
} ArcNode;

// 顶点
typedef struct VNode {
    VertexType data;  // 顶点信息
    ArcNode* first;   // 第一条边/弧
} VNode, AdjList[MaxVertexNum];

// 用邻接表（adjacency table）存储的图
typedef struct {
    AdjList vertices;  // 邻接表
    int vexnum, arcnum;
} ALGraph;

/* 图的度

无向图的度
    第 i 个结点的度 =
遍历和该顶点相关的边链表即可,有多少个边结点它的度就有多少。
    空间复杂度:同一条边实际上被存储了两次，边结点的数量是|E|，O(|V|+2|E|)

有向图的度
    第 i个结点的出度：
        需要遍历和这个结点相关的边链表，该链表反映从当前结点出去的弧。
    要找入度和指向当前结点的弧就比较麻烦。如果要统计某结点的入度和指向该结点的弧就只能遍历所有结点的边链表。

    空间复杂度:边结点的数量是|E|，O(|V|+|E|)
*/