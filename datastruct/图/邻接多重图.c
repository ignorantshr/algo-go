#define MaxVexSize 100

// 邻接多重表只适用于存储 无向图
// 空间复杂度为 O(|V|+|E|)，每条边只对应一份数据

typedef struct ENode {
    int vex1idx;  // 边的两个顶点编号
    int vex2idx;
    struct ENode* edge1;  // 依附于vex1的下一条边
    struct ENode* edge2;  // 依附于vex2的下一条边
    // int power; // 权值
} ENode;

typedef struct Node {
    int data;
    ENode* firste;  // 第一条边
} Node, List[MaxVexSize];

typedef struct LGraph {
    List vertices;
    int count, arcnum;
};
