#include "二叉树-链式存储.c"

// 链式队列结点
typedef struct LinkNode {
    BiTNode* data;  // ⚠️存指针而不是结点
    struct LinkNode* next;
} LinkNode;

typedef struct {
    LinkNode *front, *rear;  // 队头队尾
} LinkQueue;

void InitQueue(LinkQueue* q) {}
bool EnQueue(LinkQueue q, BiTNode* n) {
    return true;
}
bool DeQueue(LinkQueue q, BiTNode* n) {
    return true;
}
bool IsEmpty(LinkQueue q) {
    return true;
}

void visit(BiTNode n) {
    printf("%d ", n.data);
}

void levelBfs(BiTree t) {
    if (t == NULL) {
        return;
    }

    LinkQueue q;
    InitQueue(&q);
    EnQueue(q, t);
    while (!IsEmpty(q)) {
        BiTNode tmp;
        DeQueue(q, &tmp);
        visit(tmp);
        if (tmp.lchild != NULL) {
            EnQueue(q, tmp.lchild);
        }
        if (tmp.rchild != NULL) {
            EnQueue(q, tmp.rchild);
        }
    }
}
