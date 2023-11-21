#define MaxVexSize 100

// 十字链表只能用于存储 有向图
// 空间复杂度为 O(|V|+|E|)

typedef struct ENode {
    int tailVexidx;         // 弧尾顶点编号
    int headVexidx;         // 弧头顶点编号
    struct ENode* arcTail;  // 弧尾相同的下一条弧
    struct ENode* arcHead;  // 弧头相同的下一条弧
    // int power; // 权值
} ENode;

typedef struct VNode {
    int data;
    struct ENode* firstout;  // 该顶点作为弧尾的第一条弧
    struct ENode* firstin;   // 该顶点作为弧头的第一条弧
} VNode, AList[MaxVexSize];

typedef struct LGraph {
    AList vertices;
    int count, arcnum;
};