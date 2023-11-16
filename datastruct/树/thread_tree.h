#include "tree.h"

// 线索二叉树结点
typedef struct ThreadNode {
    ElemType data;
    struct ThreadNode *lchild, *rchild;
    struct ThreadNode* parent;  // 加上之后就变成了三叉链表
    int ltag,
        rtag;  // 左右线索标志，tag==0 表示指针指向孩子，tag==1
               // 表示指针是“线索”，左线索指向前驱结点，右线索指向后继结点
} ThreadNode, *ThreadTree;
